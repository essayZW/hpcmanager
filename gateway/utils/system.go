package utils

import (
	"context"

	"github.com/go-redis/redis/v8"
)

const installFlags = "__HPC_MANAGER_INSTALL_FLAGS__"

// IsInstall 判断系统是否已经安装初始化完成
func IsInstall(redis *redis.Client) bool {
	cmd := redis.Get(context.Background(), installFlags)
	status, err := cmd.Bool()
	if err != nil {
		return false
	}
	return status
}

// SetInstallFlag 设置系统安装flag
func SetInstallFlag(redis *redis.Client, status bool) {
	redis.Set(context.Background(), installFlags, status, 0)
}
