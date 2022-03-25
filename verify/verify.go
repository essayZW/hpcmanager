package verify

// PermissionAction 需要权限的操作
type PermissionAction string

const (
	// AddUserAction 添加用户操作
	AddUserAction = "__ADD_USER__"
	// AddUserPermissionAction 添加用户权限操作
	AddUserPermissionAction = "__ADD_USER_PERMISSION__"
	// RemoveUserPermissionAction 删除用户权限操作
	RemoveUserPermissionAction = "__REMOVE_USER_PERMISSION__"
	// AddPermission 添加权限等级
	AddPermission = "__ADD_PERMISSION__"
	// GetUserInfo 查询用户信息
	GetUserInfo = "__GET_USER_INFO__"
	// GetGroupInfo 查询用户组信息
	GetGroupInfo = "__GET_GROUP_INFO__"
	// ApplyJoinGroup 申请加入组操作
	ApplyJoinGroup = "__ADD_APPLY_GROUP__"
	// RevokeUserApplyGroup 撤销加入组的申请
	RevokeUserApplyGroup = "__REVOKE_USER_APPLY_GROUP__"
	// SearchTutorInfo 搜索导师以及组基本信息
	SearchTutorInfo = "__SEARCH_TUTOR_INFO__"
	// CheckJoinGroupApply 审核加入组申请
	CheckJoinGroupApply = "__CHECK_JOIN_GROUP_APPLY__"
	// CreateGroup 创建组
	CreateGroup = "__CREATE_GROUP__"
	// JoinGroup 用户加入某个组
	JoinGroup = "__JOIN_GROUP__"
	// CreateProject 创建新的项目记录
	CreateProject = "__CREATE_PROJECT__"
	// GetProjectInfo 查询项目信息
	GetProjectInfo = "__GET_PROJECT_INFO__"
	// CreateNodeApply 创建计算节点申请
	CreateNodeApply = "__CREATE_NODE_APPLY__"
	// RevokeNodeApply 撤销机器节点申请
	RevokeNodeApply = "__REVOKE_NODE_APPLY__"
	// GetNodeApplyInfo 查询个计算节点包机申请
	GetNodeApplyInfo = "__GET_NODE_APPLY_INFO__"
	// CheckNodeApply 审核机器节点申请
	CheckNodeApply = "__CHECK_NODE_APPLY__"
	// CreateNodeDistributeWO 创建机器节点分配工单
	CreateNodeDistributeWO = "__CREATE_NODE_DISTRIBUTE_WO__"
	// QueryNodeDistributeWO 查询机器节点分配工单
	QueryNodeDistributeWO = "__QUERY_NODE_DISTRIBUTE_WO__"
	// FinishNodeDistributeWO 处理某个机器节点分配工单
	FinishNodeDistributeWO = "__FINISH_NODE_DISTRIBUTE_WO__"
)

// Verify 进行操作的权限验证
type Verify interface {
	// Identify 权限鉴定, action 为需要鉴定的操作名称, permissionLevel 为鉴定者拥有的权限等级列表
	// 若权限验证不通过或者操作不存在则返回false,权限验证通过则返回true
	Identify(action PermissionAction, permissionLevel []Level) bool
	// AllowedActions 传入用户的权限等级，列出所有其支持的操作列表
	AllowedActions(permissionLevel []Level) []PermissionAction
}

// DefaultVerify 默认的验证器
var DefaultVerify = New()

// New 新建一个默认的Verify
func New() Verify {
	return newDefault()
}

// Identify 默认验证器进行操作权限验证
func Identify(action PermissionAction, permissionLevel []int32) bool {
	levels := make([]Level, len(permissionLevel))
	for index := range permissionLevel {
		levels[index] = Level(permissionLevel[index])
	}
	return DefaultVerify.Identify(action, levels)
}

// AllowedActions 默认验证器列出支持的操作
func AllowedActions(permissionLevel []Level) []PermissionAction {
	return DefaultVerify.AllowedActions(permissionLevel)
}

// IsAdmin 是否是管理员
func IsAdmin(permissionLevel []int32) bool {
	for _, level := range permissionLevel {
		if level == int32(CommonAdmin) || level == int32(SuperAdmin) {
			return true
		}
	}
	return false
}

// IsTutor 是否是导师
func IsTutor(permissionLevel []int32) bool {
	for _, level := range permissionLevel {
		if level == int32(Tutor) {
			return true
		}
	}
	return false
}

// Level 权限等级
type Level int

const (
	// MinLevel 权限值的下界
	MinLevel Level = -128
	// Guest 游客级别的权限，基本没有什么权限
	Guest Level = iota - 1
	// Common 普通用户的权限
	Common
	// Tutor 导师用户的权限
	Tutor
	// CommonAdmin 普通管理员的权限
	CommonAdmin
	// SuperAdmin 超级管理员的权限
	SuperAdmin
	// MaxLevel 权限值的上界
	MaxLevel Level = 128
)
