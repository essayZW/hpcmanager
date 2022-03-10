package service

import (
	"context"
	"errors"

	"github.com/essayZW/hpcmanager/db"
	"github.com/essayZW/hpcmanager/hpc/logic"
	hpcproto "github.com/essayZW/hpcmanager/hpc/proto"
	"github.com/essayZW/hpcmanager/logger"
	publicproto "github.com/essayZW/hpcmanager/proto"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/essayZW/hpcmanager/verify"
	"go-micro.dev/v4/client"
)

// HpcService hpc服务
type HpcService struct {
	hpcLogic         *logic.HpcLogic
	userService      userpb.UserService
	userGroupService userpb.GroupService
}

// Ping ping测试
func (h *HpcService) Ping(ctx context.Context, req *publicproto.Empty, resp *publicproto.PingResponse) error {
	logger.Info("Hpc PING ", resp)
	resp.Msg = "PONG"
	resp.Ip = req.BaseRequest.RequestInfo.RemoteIP
	resp.RequestId = req.BaseRequest.RequestInfo.Id
	return nil
}

// AddUserWithGroup 创建
func (h *HpcService) AddUserWithGroup(ctx context.Context, req *hpcproto.AddUserWithGroupRequest, resp *hpcproto.AddUserWithGroupResponse) error {
	logger.Infof("AddUserWithGroup: %v", req.BaseRequest)
	// 鉴权
	if !verify.Identify(verify.CreateGroup, req.BaseRequest.UserInfo.Levels) {
		logger.Info("AddUserWithGroup permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("AddUserWithGroup permission forbidden")
	}
	_, err := db.Transication(context.Background(), func(c context.Context, i ...interface{}) (interface{}, error) {
		// 通过source在机器上创建用户组并添加导师用户到新创建的组
		data, err := h.hpcLogic.AddUserWithGroup(c, req.TutorUsername, req.GroupName)
		if err != nil {
			return nil, err
		}
		groupname, ok := data["gname"].(string)
		if !ok {
			logger.Warn("AddUserWithGroup error: error response data: ", err)
			return nil, errors.New("error response data from source")
		}
		gid, ok := data["gid"].(int)
		if !ok {
			logger.Warn("AddUserWithGroup error: error response data: ", err)
			return nil, errors.New("error response data from source")
		}
		// 同步hpc_group表中的相关信息
		hpcGroupID, err := h.hpcLogic.CreateGroup(c, groupname, req.QueueName, gid)
		if err != nil {
			// NOTE 数据库数据更新失败,但是执行的脚本命令不可以恢复,需要联系管理员手动同步数据
			return nil, err
		}
		resp.HpcGroupID = int32(hpcGroupID)
		resp.GroupName = groupname
		resp.Gid = int32(gid)
		halfFlag := false
		username, ok := data["uname"].(string)
		if !ok {
			halfFlag = true
		}
		uid, ok := data["uid"].(int)
		if !ok {
			halfFlag = true
		}
		if halfFlag {
			// 只成功创建了组，没有成功创建用户
			// NOTE 这种情况应该认为操作不成功,但是返回err会导致group回滚从而导致数据不一致
			// 因此可能需要联系管理员进行手动数据同步
			return nil, nil
		}
		// 同步hpc_user表中的相关信息
		hpcUserID, err := h.hpcLogic.CreateUser(c, username, uid)
		if err != nil {
			// NOTE 数据库数据更新失败,但是执行的脚本命令不可以恢复,需要联系管理员手动同步数据
			return nil, err
		}
		resp.HpcUserID = int32(hpcUserID)
		resp.UserName = username
		resp.Uid = int32(uid)
		return nil, nil
	})
	return err
}

// AddUserToGroup 添加用户到用户组
func (h *HpcService) AddUserToGroup(ctx context.Context, req *hpcproto.AddUserToGroupRequest, resp *hpcproto.AddUserToGroupResponse) error {
	logger.Infof("AddUserToGroup : %v", req.BaseRequest)
	// 鉴权
	if !verify.Identify(verify.AddUserAction, req.BaseRequest.UserInfo.Levels) {
		logger.Info("AddUserToGroup permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("AddUserToGroup permission forbidden")
	}
	_, err := db.Transication(context.Background(), func(c context.Context, i ...interface{}) (interface{}, error) {
		groupInfo, err := h.hpcLogic.GetGroupInfoByID(c, int(req.HpcGroupID))
		if err != nil {
			return nil, errors.New("invalid hpc group id")
		}
		data, err := h.hpcLogic.AddUserToGroup(c, req.UserName, groupInfo.Name, groupInfo.GID)
		if err != nil {
			return nil, err
		}
		username, ok := data["uname"].(string)
		if !ok {
			logger.Warn("AddUserToGroup error: error response data: ", err)
			return nil, errors.New("error response data from source")
		}
		uid, ok := data["uid"].(int)
		if !ok {
			logger.Warn("AddUserToGroup error: error response data: ", err)
			return nil, errors.New("error response data from source")
		}
		userID, err := h.hpcLogic.CreateUser(c, username, uid)
		if err != nil {
			return nil, err
		}
		resp.HpcUserID = int32(userID)
		resp.Uid = int32(uid)
		resp.UserName = username
		return nil, nil
	})
	return err
}

// GetUserInfoByID 通过ID查询hpc_user的信息
func (h *HpcService) GetUserInfoByID(ctx context.Context, req *hpcproto.GetUserInfoByIDRequest, resp *hpcproto.GetUserInfoByIDResponse) error {
	logger.Infof("GetUserInfoByID: %v", req.BaseRequest)
	if !verify.Identify(verify.GetUserInfo, req.BaseRequest.UserInfo.Levels) {
		logger.Info("GetUserInfoByID permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("GetUserInfoByID permission forbidden")
	}
	info, err := h.hpcLogic.GetUserInfoByID(context.Background(), int(req.HpcUserID))
	if err != nil {
		return errors.New("query hpc_user info error")
	}
	isAdmin := verify.IsAdmin(req.BaseRequest.UserInfo.Levels)
	isTutor := verify.IsTutor(req.BaseRequest.UserInfo.Levels)
	if !isTutor && !isAdmin {
		// 普通用户,需判断自己是不是该hpc_user的记录对应者
		userResp, err := h.userService.GetUserInfo(ctx, &userpb.GetUserInfoRequest{
			Userid:      req.BaseRequest.UserInfo.UserId,
			BaseRequest: req.BaseRequest,
		})
		if err != nil {
			return errors.New("user info get error")
		}
		if userResp.UserInfo.HpcUserID != req.HpcUserID {
			return errors.New("user only can query self hpc user info")
		}
	} else if !isAdmin && isTutor {
		// 导师用户,需判断该hpc_user对应的用户是否属于自己的组
		userResp, err := h.userService.GetUserInfoByHpcID(ctx, &userpb.GetUserInfoByHpcIDRequest{
			BaseRequest: req.BaseRequest,
			HpcUserID:   req.HpcUserID,
		})
		if err != nil {
			return errors.New("user info get error")
		}
		if userResp.Info.GroupId != req.BaseRequest.UserInfo.GroupId {
			return errors.New("tutor only can query self group's user info")
		}
	}
	resp.User = &hpcproto.HpcUser{
		Id:             int32(info.ID),
		NodeUsername:   info.NodeUsername,
		NodeUID:        int32(info.NodeUID),
		NodeMaxQuota:   int32(info.NodeMaxQuota),
		QuotaStartTime: info.QuotaStartTime.Time.Unix(),
		QuotaEndTime:   info.QuotaEndTime.Time.Unix(),
	}
	if info.ExtraAttributes != nil {
		resp.User.ExtraAttributes = info.ExtraAttributes.String()
	}
	return nil
}

// GetGroupInfoByID 通过id查询hpc_group的信息
func (h *HpcService) GetGroupInfoByID(ctx context.Context, req *hpcproto.GetGroupInfoByIDRequest, resp *hpcproto.GetGroupInfoByIDResponse) error {
	logger.Infof("GetGroupInfoByID: %v", req.BaseRequest)
	if !verify.Identify(verify.GetGroupInfo, req.BaseRequest.UserInfo.Levels) {
		logger.Info("GetGroupInfoByID permission forbidden: ", req.BaseRequest.RequestInfo.Id, ", fromUserId: ", req.BaseRequest.UserInfo.UserId, ", withLevels: ", req.BaseRequest.UserInfo.Levels)
		return errors.New("GetGroupInfoByID permission forbidden")
	}

	isAdmin := verify.IsAdmin(req.BaseRequest.UserInfo.Levels)
	isTutor := verify.IsTutor(req.BaseRequest.UserInfo.Levels)
	if !isAdmin && isTutor {
		// 是导师用户,需要查询导师用户所属组是不是当前查询的组
		groupResp, err := h.userGroupService.GetGroupInfoByID(ctx, &userpb.GetGroupInfoByIDRequest{
			GroupID:     req.BaseRequest.UserInfo.GroupId,
			BaseRequest: req.BaseRequest,
		})
		if err != nil {
			return errors.New("get group info error")
		}
		if groupResp.GroupInfo.HpcGroupID != req.HpcGroupID {
			return errors.New("get group info error: permision forbiden")
		}
	}
	info, err := h.hpcLogic.GetGroupInfoByID(context.Background(), int(req.HpcGroupID))
	if err != nil {
		return errors.New("hpc group info query error")
	}

	resp.Group = &hpcproto.HpcGroup{
		Id:        int32(info.ID),
		Name:      info.Name,
		QueueName: info.QueueName,
		GID:       int32(info.GID),
	}
	if info.ExtraAttributes != nil {
		resp.Group.ExtraAttributes = info.ExtraAttributes.String()
	}
	return nil
}

var _ hpcproto.HpcHandler = (*HpcService)(nil)

// NewHpc 新建一个Hpc服务
func NewHpc(client client.Client, hpcLogic *logic.HpcLogic) *HpcService {
	return &HpcService{
		hpcLogic:         hpcLogic,
		userService:      userpb.NewUserService("user", client),
		userGroupService: userpb.NewGroupService("user", client),
	}
}
