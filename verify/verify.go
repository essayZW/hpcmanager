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
