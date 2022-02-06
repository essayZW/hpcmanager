package hpcmanager

import "flag"

var (
	etcdAddress string
)

// LoadCommonArgs 注册通用的命令行参数并解析到变量上
func LoadCommonArgs() {
	flag.StringVar(&etcdAddress, "etcdAddress", "172.17.0.2:2379", "etcd service address")
}

// GetEtcdAddress 获得传入的etcd服务地址
func GetEtcdAddress() string {
	return etcdAddress
}
