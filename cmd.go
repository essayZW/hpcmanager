package hpcmanager

import "flag"

var (
	etcdAddress string
	hasLoad     bool = false
)

const (
	// EnvName 应用环境的环境变量名
	EnvName = "HPCMANAGER_ENV"
	// ProductionEnvValue 生产环境的环境变量值
	ProductionEnvValue = "production"
	// DevelopmentEnvValue 开发环境的环境变量值
	DevelopmentEnvValue = "development"
)

// check 检查是否已经调用过LoadCommonArgs
func check() {
	if !hasLoad {
		panic("Must call LoadCommonArgs before get common cmd args")
	}
}

// LoadCommonArgs 注册通用的命令行参数并解析到变量上
func LoadCommonArgs() {
	flag.StringVar(&etcdAddress, "etcdAddress", "172.17.0.2:2379", "etcd service address")
	hasLoad = true
}

// GetEtcdAddress 获得传入的etcd服务地址
func GetEtcdAddress() string {
	check()
	return etcdAddress
}
