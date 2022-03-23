package source

import "go-micro.dev/v4/logger"

// HpcSource hpc调度系统的操作接口定义
// 当前子项目用于hpc作业调度系统集成到微服务系统中的一个包装
// 考虑到可能有不同形式的作业调度系统接口,因此定义统一的操作接口
type HpcSource interface {
	// AddUserWithGroup 创建一个用户组并添加用户到该用户组中,返回用户ID、组ID以及用户名和组名
	AddUserWithGroup(userName string, groupName string) (map[string]interface{}, error)
	// AddUserToGroup 添加用户到现有的用户组
	AddUserToGroup(userName string, groupName string, gid int) (map[string]interface{}, error)
}

// New 创建默认的作业调度源
func New(options ...Option) HpcSource {
	opts := Options{}
	for _, opt := range options {
		opt(&opts)
	}
	if opts.DevMode {
		logger.Info("Use devMode source")
		return newDev(&opts)
	}
	return newSource(&opts)
}
