package controller

import (
	"context"
	"time"

	"github.com/essayZW/hpcmanager/gateway/middleware"
	gatewaypb "github.com/essayZW/hpcmanager/gateway/proto"
	"github.com/essayZW/hpcmanager/gateway/request/json"
	"github.com/essayZW/hpcmanager/gateway/response"
	"github.com/essayZW/hpcmanager/gateway/utils"
	"github.com/essayZW/hpcmanager/logger"
	permissionpb "github.com/essayZW/hpcmanager/permission/proto"
	userpb "github.com/essayZW/hpcmanager/user/proto"
	"github.com/essayZW/hpcmanager/verify"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go-micro.dev/v4/client"
)

// System 系统相关功能的控制器
type System struct {
	userService       userpb.UserService
	permissionService permissionpb.PermissionService
	redisConn         *redis.Client
}

// install /api/sys/install POST 初始化系统相关的设置
func (sys *System) install(ctx *gin.Context) {
	var params json.CreateUserParam
	if err := ctx.ShouldBindJSON(&params); err != nil {
		rep := response.New(500, err.Error(), false, "invalid user params")
		rep.Send(ctx)
		return
	}
	baseReq, _ := ctx.Get(middleware.BaseRequestKey)
	baseRequest := baseReq.(*gatewaypb.BaseRequest)
	// 由于初始化系统需要的部分操作需要管理员权限,因此为此次操作赋予临时管理员权限
	baseRequest.UserInfo.Levels = append(baseRequest.UserInfo.Levels, int32(verify.SuperAdmin))
	// 1. 检查是否已经安装
	if utils.IsInstall(sys.redisConn) {
		resp := response.New(200, nil, false, "系统已经安装")
		resp.Send(ctx)
		return
	}
	// 2. 初始化系统权限等级表
	c, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	if !sys.initPermision(c, baseRequest) {
		resp := response.New(500, nil, false, "系统权限安装失败")
		resp.Send(ctx)
		return
	}
	// 3. 创建SuperAdmin权限用户
	userResp, err := sys.userService.AddUser(c, &userpb.AddUserRequest{
		UserInfo: &userpb.UserInfo{
			Username: params.Username,
			Password: params.Password,
			Name:     params.Name,
			College:  params.CollegeName,
			Email:    params.Email,
			Tel:      params.Tel,
		},
		BaseRequest: baseRequest,
	})
	if err != nil {
		logger.Warn("add user error: ", err)
		resp := response.New(500, nil, false, "管理员用户添加失败")
		resp.Send(ctx)
		return
	}
	resp, err := sys.permissionService.AddUserPermission(c, &permissionpb.AddUserPermissionRequest{
		Userid:      userResp.Userid,
		Level:       int32(verify.SuperAdmin),
		BaseRequest: baseRequest,
	})
	if err != nil || !resp.Success {
		logger.Warn("add super admin permission error: ", err)
		resp := response.New(500, nil, false, "用户权限初始化失败")
		resp.Send(ctx)
		return
	}
	// 4. 重置安装flag
	utils.SetInstallFlag(sys.redisConn, true)
	response := response.New(200, nil, true, "success")
	response.Send(ctx)
}

func (sys *System) initPermision(ctx context.Context, baseReq *gatewaypb.BaseRequest) bool {

	permissions := []struct {
		Name        string
		Level       verify.Level
		Description string
	}{
		{"Guest", verify.Guest, "Guest一般赋予没有加入任何组的新用户,代表基本没有权限"},
		{"Common", verify.Common, "Common一般指普通用户,也就是学生用户权限"},
		{"Tutor", verify.Tutor, "Tutor一般指导师用户权限"},
		{"CommonAdmin", verify.CommonAdmin, "CommonAdmin一般指系统管理员,具有大部分权限"},
		{"SuperAdmin", verify.SuperAdmin, "SuperAdmin一般指超级管理员,可以任命普通管理员"},
	}
	for index := range permissions {
		addResp, err := sys.permissionService.AddPermission(ctx, &permissionpb.AddPermissionRequest{
			Info: &permissionpb.PermissionInfo{
				Name:        permissions[index].Name,
				Level:       int32(permissions[index].Level),
				Description: permissions[index].Description,
			},
			BaseRequest: baseReq,
		})
		if err != nil {
			logger.Warn("add permission error: ", err)
			return false
		}
		if addResp.PermissionID == 0 {
			return false
		}
	}
	return true
}

// Registry 注册system控制器的相关函数
func (sys *System) Registry(router *gin.RouterGroup) *gin.RouterGroup {
	logger.Info("registry gateway controller System")
	sysRouter := router.Group("/sys")

	sysRouter.POST("/install", sys.install)
	middleware.RegistryExcludeAPIPath("POST:/api/sys/install")
	return sysRouter
}

// NewSystem 创建新的系统功能控制器
func NewSystem(client client.Client, redisConn *redis.Client) Controller {
	userService := userpb.NewUserService("user", client)
	permissionService := permissionpb.NewPermissionService("permission", client)
	return &System{
		userService:       userService,
		permissionService: permissionService,
		redisConn:         redisConn,
	}

}
