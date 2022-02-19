package verify

// hardcodeVerify 权限定义硬编码的权限验证器，为默认的权限验证器
type hardcodeVerify struct {
	actionsLevel map[PermissionAction]*actionVerify
}

func (verify *hardcodeVerify) Identify(action PermissionAction, levels []Level) bool {
	if av, ok := verify.actionsLevel[action]; ok {
		return av.identify(levels)
	}
	return false
}

func (verify *hardcodeVerify) AllowedActions(permissionLevel []Level) []PermissionAction {
	res := make([]PermissionAction, 0)
	for k, v := range verify.actionsLevel {
		if v.identify(permissionLevel) {
			res = append(res, k)
		}
	}
	return res
}

func newDefault() *hardcodeVerify {
	return &hardcodeVerify{
		actionsLevel: map[PermissionAction]*actionVerify{
			// 普通管理员及以上
			AddUserAction: {
				maxLevel: MaxLevel,
				minLevel: CommonAdmin,
			},
			// 普通管理员及以上
			AddUserPermissionAction: {
				maxLevel: MaxLevel,
				minLevel: CommonAdmin,
			},
			// 普通管理员及以上
			RemoveUserPermissionAction: {
				maxLevel: MaxLevel,
				minLevel: CommonAdmin,
			},
			// 超级管理员及以上
			AddPermission: {
				maxLevel: MaxLevel,
				minLevel: SuperAdmin,
			},
			// 普通成员权限及以上
			GetUserInfo: {
				maxLevel: MaxLevel,
				minLevel: Common,
			},
			// 导师权限及以上
			GetGroupInfo: {
				maxLevel: MaxLevel,
				minLevel: Tutor,
			},
			// 游客权限
			ApplyJoinGroup: {
				maxLevel: Guest,
				minLevel: Guest,
			},
			// 游客权限
			SearchTutorInfo: {
				maxLevel: Guest,
				minLevel: Guest,
			},
		},
	}
}

// actionVerify PermissionAction的加强版本，包含了对于单个操作的具体权限要求的定义及其权限验证函数
type actionVerify struct {
	maxLevel Level
	minLevel Level
}

func (av *actionVerify) identify(levels []Level) bool {
	for _, level := range levels {
		if level >= av.minLevel && level <= av.maxLevel {
			return true
		}
	}
	return false
}
