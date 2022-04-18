package source

import (
	"context"
	"time"

	"go-micro.dev/v4/logger"
)

// HpcSource hpc调度系统的操作接口定义
// 当前子项目用于hpc作业调度系统集成到微服务系统中的一个包装
// 考虑到可能有不同形式的作业调度系统接口,因此定义统一的操作接口
type HpcSource interface {
	// AddUserWithGroup 创建一个用户组并添加用户到该用户组中,返回用户ID、组ID以及用户名和组名
	AddUserWithGroup(userName string, groupName string) (map[string]interface{}, error)
	// AddUserToGroup 添加用户到现有的用户组
	AddUserToGroup(userName string, groupName string, gid int) (map[string]interface{}, error)
	// GetNodeUsageWithDate 获取某段时间内的节点使用情况
	GetNodeUsageWithDate(ctx context.Context, startTime, endTime time.Time) ([]*HpcNodeUsage, error)
	// QuotaQuery 用户的存储信息查询
	QuotaQuery(username string, fs string) (*QuotaUsageInfo, error)
}

// New 创建默认的作业调度源
func New(options ...Option) (HpcSource, error) {
	opts := Options{}
	for _, opt := range options {
		opt(&opts)
	}
	if opts.DevMode {
		logger.Info("Use devMode source")
		return newDev(&opts), nil
	}
	return newSource(&opts)
}

// HpcNodeUsage 计算节点的使用情况记录的映射
type HpcNodeUsage struct {
	Username  string  `db:"UserName"`
	GroupName string  `db:"GroupName"`
	QueueName string  `db:"Queue"`
	WallTime  float64 `db:"WallTime"`
	GWallTime float64 `db:"GpusWallTime"`
}

// QuotaUsageInfo 存储使用信息
// NOTE: 作业调度系统中返回的存储数据不是纯数字数据，是带单位的容量数据
// 比如0k, 1k, 1T这种
type QuotaUsageInfo struct {
	Used string
	Max  string
}
