package source

import (
	"github.com/essayZW/hpcmanager/config"
	"github.com/go-redis/redis/v8"
)

// Options source配置的options
type Options struct {
	// CmdLocation 脚本文件的根目录
	CmdBaseDir string

	// DevMode 是否是开发模式下
	DevMode bool

	// dbConf 数据库配置
	dbConf *config.Database

	// redisConn redis连接
	redisConn *redis.Client
}

// Option 选项
type Option func(*Options)

// WithCmdBaseDir 配置可选的脚本文件路径
func WithCmdBaseDir(dir string) Option {
	return func(o *Options) {
		o.CmdBaseDir = dir
	}
}

// WithDevSource 配置是否使用dev模式下的source
func WithDevSource(dev bool) Option {
	return func(o *Options) {
		o.DevMode = dev
	}
}

// WithDBSource 配置作业调度系统的数据库
func WithDBSource(dbConf *config.Database) Option {
	return func(o *Options) {
		o.dbConf = dbConf
	}
}

// WithDevRedis 配置dev模式下的redis数据库链接
func WithDevRedis(redisConn *redis.Client) Option {
	return func(o *Options) {
		o.redisConn = redisConn
	}
}
