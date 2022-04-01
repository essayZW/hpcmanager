package main

import (
	"context"
	"time"

	gatewaypb "github.com/essayZW/hpcmanager/gateway/proto"
	hpcpb "github.com/essayZW/hpcmanager/hpc/proto"
	"github.com/essayZW/hpcmanager/logger"
	nodepb "github.com/essayZW/hpcmanager/node/proto"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/essayZW/hpcmanager/verify"
	"go-micro.dev/v4/client"
)

// RPCLoader 通过RPC协议进行数据的获取以及同步
type RPCLoader struct {
	nodeService      nodepb.NodeService
	hpcService       hpcpb.HpcService
	userService      userpb.UserService
	userGroupService userpb.GroupService

	maxRetryCount int
}

func (loader *RPCLoader) generateBaseRequest() *gatewaypb.BaseRequest {
	return &gatewaypb.BaseRequest{
		UserInfo: &gatewaypb.UserInfo{
			Levels: []int32{
				int32(verify.SuperAdmin),
			},
		},
		RequestInfo: &gatewaypb.RequestInfo{
			Id: "__SYSTEM_CRONTAB__",
		},
	}
}

type cacheItem struct {
	userID        int
	username      string
	userName      string
	tutorID       int
	tutorUsername string
	tutorName     string
}

// Sync 从hpc服务同步数据到node服务
func (loader *RPCLoader) Sync(ctx context.Context, startDate, endDate time.Time) error {
	baseRequest := loader.generateBaseRequest()
	infos, err := loader.hpcService.GetNodeUsage(ctx, &hpcpb.GetNodeUsageRequest{
		BaseRequest:   baseRequest,
		StartTimeUnix: startDate.Unix(),
		EndTimeUnix:   endDate.Unix(),
	})
	if err != nil {
		return err
	}

	cacheMap := make(map[string]cacheItem)
	for _, usage := range infos.Usages {
		handler := func() error {
			var cacheInfo cacheItem
			var ok bool
			cacheInfo, ok = cacheMap[usage.Username]
			if !ok {
				// 查询缓存数据
				// 首先查询hpc信息
				hpcUserResp, err := loader.hpcService.GetUserInfoByUsername(ctx, &hpcpb.GetUserInfoByUsernameRequest{
					BaseRequest: baseRequest,
					Username:    usage.Username,
				})
				if err != nil {
					return err
				}
				// 通过hpc信息查询用户信息
				userInfoResp, err := loader.userService.GetUserInfoByHpcID(ctx, &userpb.GetUserInfoByHpcIDRequest{
					BaseRequest: baseRequest,
					HpcUserID:   hpcUserResp.User.Id,
				})
				if err != nil {
					return err
				}
				// 查询hpc group信息
				hpcGroupResp, err := loader.hpcService.GetGroupInfoByGroupName(ctx, &hpcpb.GetGroupInfoByGroupNameRequest{
					BaseRequest: baseRequest,
					Name:        usage.GroupName,
				})
				if err != nil {
					return err
				}
				groupResp, err := loader.userGroupService.GetGroupInfoByHpcID(ctx, &userpb.GetGroupInfoByHpcIDRequest{
					BaseRequest: baseRequest,
					HpcGroupID:  hpcGroupResp.Group.Id,
				})
				cacheInfo = cacheItem{
					userID:        int(userInfoResp.Info.Id),
					username:      userInfoResp.Info.Username,
					userName:      userInfoResp.Info.Name,
					tutorID:       int(groupResp.GroupInfo.TutorID),
					tutorUsername: groupResp.GroupInfo.TutorUsername,
					tutorName:     groupResp.GroupInfo.TutorName,
				}
				logger.Debug("load cache item: ", usage.Username)
				cacheMap[usage.Username] = cacheInfo
			} else {
				logger.Debug("use cache item: ", usage.Username)
			}
			resp, err := loader.nodeService.AddNodeUsageTimeRecord(ctx, &nodepb.AddNodeUsageTimeRecordRequest{
				BaseRequest:   baseRequest,
				UserID:        int32(cacheInfo.userID),
				Username:      cacheInfo.username,
				Name:          cacheInfo.userName,
				HpcUserName:   usage.Username,
				TutorID:       int32(cacheInfo.tutorID),
				TutorUsername: cacheInfo.tutorUsername,
				TutorName:     cacheInfo.tutorName,
				HpcGroupName:  usage.GroupName,
				QueueName:     usage.QueueName,
				WallTime:      usage.WallTime,
				GwallTime:     usage.GwallTime,
				StartTimeUnix: startDate.Unix(),
				EndTimeUnix:   endDate.Unix(),
			})
			if err != nil {
				delete(cacheMap, usage.Username)
				return err
			}
			logger.Info("loader sync a info: ", resp.Id)
			return nil
		}

		var retryErr error
		for i := 0; i < loader.maxRetryCount; i++ {
			// 进行多次重试
			retryErr = handler()
			if retryErr == nil {
				break
			} else {
				logger.Warn(retryErr)
			}
		}
		if retryErr != nil {
			// 重试失败
			logger.Errorf("loader sync error with usage record: %v, with error: %v", usage, retryErr)
		}
	}
	return nil
}

// NewLoader 创建新的机器时长数据的加载器
func NewLoader(client client.Client, retryCount int) *RPCLoader {
	nodeService := nodepb.NewNodeService("node", client)
	hpcService := hpcpb.NewHpcService("hpc", client)
	userService := userpb.NewUserService("user", client)
	groupService := userpb.NewGroupService("user", client)
	return &RPCLoader{
		nodeService:      nodeService,
		hpcService:       hpcService,
		userService:      userService,
		userGroupService: groupService,
		maxRetryCount:    retryCount,
	}
}
