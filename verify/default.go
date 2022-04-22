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
			// 超级管理员及以上
			AddUserPermissionAction: {
				maxLevel: MaxLevel,
				minLevel: SuperAdmin,
			},
			// 超级管理员及以上
			RemoveUserPermissionAction: {
				maxLevel: MaxLevel,
				minLevel: SuperAdmin,
			},
			// 超级管理员及以上
			AddPermission: {
				maxLevel: MaxLevel,
				minLevel: SuperAdmin,
			},
			// 游客权限及以上
			GetUserInfo: {
				maxLevel: MaxLevel,
				minLevel: Guest,
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
			RevokeUserApplyGroup: {
				maxLevel: Guest,
				minLevel: Guest,
			},
			// 游客权限
			SearchTutorInfo: {
				maxLevel: Guest,
				minLevel: Guest,
			},
			// 需要导师及以上权限
			CheckJoinGroupApply: {
				maxLevel: MaxLevel,
				minLevel: Tutor,
			},
			// 需要普通管理员及以上权限
			CreateGroup: {
				maxLevel: MaxLevel,
				minLevel: CommonAdmin,
			},
			// 需要普通管理员及以上权限
			JoinGroup: {
				maxLevel: MaxLevel,
				minLevel: CommonAdmin,
			},
			// 暂定所有权限等级都可以
			CreateProject: {
				maxLevel: MaxLevel,
				minLevel: Guest,
			},
			// 暂定所有权限等级都可以
			GetProjectInfo: {
				maxLevel: MaxLevel,
				minLevel: Guest,
			},
			// 需要普通用户及以上权限
			CreateNodeApply: {
				maxLevel: MaxLevel,
				minLevel: Common,
			},
			// 需要普通用户及以上权限
			RevokeNodeApply: {
				maxLevel: MaxLevel,
				minLevel: Common,
			},
			// 需要普通用户及以上权限
			GetNodeApplyInfo: {
				maxLevel: MaxLevel,
				minLevel: Common,
			},
			UpdateNodeApply: {
				maxLevel: MaxLevel,
				minLevel: Common,
			},
			// 需要导师及以上的权限
			CheckNodeApply: {
				maxLevel: MaxLevel,
				minLevel: Tutor,
			},
			// 暂定需要普通管理员及以上的权限
			CreateNodeDistributeWO: {
				maxLevel: MaxLevel,
				minLevel: CommonAdmin,
			},
			// 需要普通管理员及以上的权限
			QueryNodeDistributeWO: {
				maxLevel: MaxLevel,
				minLevel: CommonAdmin,
			},
			// 需要普通管理员及以上权限
			FinishNodeDistributeWO: {
				maxLevel: MaxLevel,
				minLevel: CommonAdmin,
			},
			// 需要普通用户及以上权限
			QueryNodeUsage: {
				maxLevel: MaxLevel,
				minLevel: Common,
			},
			// 需要超级管理员及以上权限
			AddNodeUsage: {
				maxLevel: MaxLevel,
				minLevel: SuperAdmin,
			},
			// 需要管理员以上权限
			CreateNodeDistributeBill: {
				maxLevel: MaxLevel,
				minLevel: CommonAdmin,
			},
			// 需要普通权限以上
			QueryNodeDistributeBill: {
				maxLevel: MaxLevel,
				minLevel: Common,
			},
			// 需要普通管理员及以上权限
			AddGroupBalance: {
				maxLevel: MaxLevel,
				minLevel: CommonAdmin,
			},
			// 需要普通管理员及以上权限
			PayNodeDistributeBill: {
				maxLevel: MaxLevel,
				minLevel: CommonAdmin,
			},
			// 需要普通管理员及以上权限
			CreateNodeWeekUsageBill: {
				maxLevel: MaxLevel,
				minLevel: CommonAdmin,
			},
			// 需要普通权限以上
			QueryNodeWeekUsageBill: {
				maxLevel: MaxLevel,
				minLevel: Common,
			},
			// 需要普通管理员及以上权限
			PayNodeUsageBill: {
				maxLevel: MaxLevel,
				minLevel: CommonAdmin,
			},
			// 需要普通权限以上
			QueryUserHpcQuota: {
				maxLevel: MaxLevel,
				minLevel: Common,
			},
			// 需要普通管理员及以上权限
			UpdateUserHpcQuota: {
				maxLevel: MaxLevel,
				minLevel: CommonAdmin,
			},
			// 需要普通管理员及以上权限
			CreateNodeQuotaModifyBill: {
				maxLevel: MaxLevel,
				minLevel: CommonAdmin,
			},
			// 需要普通权限以上
			QueryNodeQuotaBill: {
				maxLevel: MaxLevel,
				minLevel: Common,
			},
			// 需要普通管理员及以上权限
			PayNodeQuotaBill: {
				maxLevel: MaxLevel,
				minLevel: CommonAdmin,
			},
			// 需要普通权限以上
			StoreFile: {
				maxLevel: MaxLevel,
				minLevel: Common,
			},
			// 需要普通权限以上
			CreatePaperAward: {
				maxLevel: MaxLevel,
				minLevel: Common,
			},
			// 需要普通权限以上
			QueryPaperAwardApply: {
				maxLevel: MaxLevel,
				minLevel: Common,
			},
			// 需要普通管理员及以上权限
			CheckPaperAwardApply: {
				maxLevel: MaxLevel,
				minLevel: CommonAdmin,
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
