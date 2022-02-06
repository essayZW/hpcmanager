package logger

import (
	"fmt"
	"os"

	z "github.com/asim/go-micro/plugins/logger/zap/v4"
	"go-micro.dev/v4/logger"
	"go.uber.org/zap"
)

// New 创建一个新的logger
// 日志信息将会被输出到~/log/hpcmanager/{name}目录下
// 其中name是服务名,不同的服务的日志文件将会隔开
// 日志文件按照日期进行分片,文件名为日期信息
// 日志的实现为zap库
func New(name string) (logger.Logger, error) {
	var zapConfig zap.Config
	var production string
	if production = os.Getenv("PRODUCTION"); production != "" {
		// 如果PRODUCTION环境变量存在则代表为生产环境
		zapConfig = zap.NewProductionConfig()
	} else {
		zapConfig = zap.NewDevelopmentConfig()
	}
	homedir, _ := os.UserHomeDir()
	logDir := fmt.Sprintf("%s/log/hpcmanager/%s", homedir, name)
	// 判断日志目录是否存在
	if !pathExists(logDir) {
		err := os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	// 添加日志输出到文件
	zapConfig.OutputPaths = append(zapConfig.OutputPaths, logDir+"/log.txt")
	// 需要通过Options设置日志等级
	if production == "" {
		return z.NewLogger(
			z.WithConfig(zapConfig),
			logger.WithLevel(logger.DebugLevel),
		)
	}
	return z.NewLogger(
		z.WithConfig(zapConfig),
	)
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
