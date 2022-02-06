package logger

import (
	"fmt"
	"os"
	"sync"
	"time"

	z "github.com/asim/go-micro/plugins/logger/zap/v4"
	"go-micro.dev/v4/logger"
	"go.uber.org/zap"
)

var (
	cachedLogger logger.Logger
	cachedError  error
	createDate   string
	mutex        sync.Mutex
)

// New 创建一个新的logger
// 日志信息将会被输出到~/log/hpcmanager/{name}目录下
// 其中name是服务名,不同的服务的日志文件将会隔开
// 日志文件按照日期进行分片,文件名为日期信息
// 日志的实现为zap库
func New(name string) (logger.Logger, error) {
	mutex.Lock()
	defer mutex.Unlock()

	// 判断原来的logger是否已经过期
	today := time.Now().Format("2006-01-02")
	if today == createDate {
		return cachedLogger, cachedError
	}
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
	filename := fmt.Sprintf("/log-%s.txt", today)
	zapConfig.OutputPaths = append(zapConfig.OutputPaths, logDir+filename)
	// 需要通过Options设置日志等级
	if production == "" {
		cachedLogger, cachedError = z.NewLogger(
			z.WithConfig(zapConfig),
			logger.WithLevel(logger.DebugLevel),
		)
	} else {
		cachedLogger, cachedError = z.NewLogger(
			z.WithConfig(zapConfig),
		)
	}
	createDate = today
	return cachedLogger, cachedError
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

var serviceName string = "log"

// SetName 设置输出日志的服务名称
func SetName(name string) {
	serviceName = name
}

// Debug debug日志
func Debug(msg ...interface{}) {
	log, _ := New(serviceName)
	log.Log(logger.DebugLevel, msg...)
}

// Debugf 格式化输出debug日志
func Debugf(format string, v ...interface{}) {
	log, _ := New(serviceName)
	log.Logf(logger.DebugLevel, format, v...)
}

// Info info日志
func Info(msg ...interface{}) {
	log, _ := New(serviceName)
	log.Log(logger.InfoLevel, msg...)
}

// Infof 格式化输出debug日志
func Infof(format string, v ...interface{}) {
	log, _ := New(serviceName)
	log.Logf(logger.InfoLevel, format, v...)
}

// Warn warn日志
func Warn(msg ...interface{}) {
	log, _ := New(serviceName)
	log.Log(logger.WarnLevel, msg...)
}

// Warnf 格式化输出Warn日志
func Warnf(format string, v ...interface{}) {
	log, _ := New(serviceName)
	log.Logf(logger.WarnLevel, format, v...)
}

// Error error 日志
func Error(msg ...interface{}) {
	log, _ := New(serviceName)
	log.Log(logger.ErrorLevel, msg...)
}

// Errorf 格式化输出error日志
func Errorf(format string, v ...interface{}) {
	log, _ := New(serviceName)
	log.Logf(logger.ErrorLevel, format, v...)
}

// Fatal fatal日志
func Fatal(msg ...interface{}) {
	log, _ := New(serviceName)
	log.Log(logger.FatalLevel, msg...)
}

// Fatalf 格式化输出fatal日志
func Fatalf(format string, v ...interface{}) {
	log, _ := New(serviceName)
	log.Logf(logger.FatalLevel, format, v...)
}
