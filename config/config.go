package config

import (
	"fmt"
	"os"

	yaml "github.com/asim/go-micro/plugins/config/encoder/yaml/v4"
	"github.com/essayZW/hpcmanager"
	"github.com/essayZW/hpcmanager/logger"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source/file"
)

var (
	configFileDir string
)

const (
	configFileBaseName = "config"
	configFileSuffix   = "yaml"
)

func init() {
	osConfigDir, err := os.UserConfigDir()
	if err != nil {
		osConfigDir = "/config"
	}
	configFileDir = osConfigDir + "/hpcmanager"
}

// LoadConfigSource 加载配置源
func LoadConfigSource() (config.Config, error) {
	enc := yaml.NewEncoder()
	c, _ := config.NewConfig(
		config.WithReader(
			json.NewReader(
				reader.WithEncoder(enc),
			),
		),
	)
	filePath := fmt.Sprintf("%s/%s-%s.%s", configFileDir, configFileBaseName, getEnvValue(), configFileSuffix)
	logger.Info("load config file from path ", filePath)
	// load the config from a file source
	if err := c.Load(file.NewSource(
		file.WithPath(filePath),
	)); err != nil {
		logger.Error(err)
		return nil, err
	}
	return c, nil

}

// LoadDatabase 加载数据库配置
func LoadDatabase() (*Database, error) {
	c, err := LoadConfigSource()
	if err != nil {
		return nil, err
	}

	var database Database

	if err := c.Get("database").Scan(&database); err != nil {
		return nil, err
	}
	return &database, nil
}

func getEnvValue() string {
	value := os.Getenv(hpcmanager.EnvName)
	if value != hpcmanager.ProductionEnvValue {
		if value == "" {
			return hpcmanager.DevelopmentEnvValue
		}
		return value
	}
	return hpcmanager.ProductionEnvValue
}

// LoadRedis 加载redis服务配置
func LoadRedis() (*Redis, error) {
	c, err := LoadConfigSource()
	if err != nil {
		return nil, err
	}
	var redis Redis
	if err := c.Get("redis").Scan(&redis); err != nil {
		return nil, err
	}
	return &redis, err
}

// LoadRegistry 加载etcd注册中心配置
func LoadRegistry() (*Registry, error) {
	c, err := LoadConfigSource()
	if err != nil {
		return nil, err
	}
	var registry Registry
	if err := c.Get("registry").Scan(&registry); err != nil {
		return nil, err
	}
	return &registry, err
}

// LoadRabbitmq 加载rabbitmq消息队列配置
func LoadRabbitmq() (*Rabbitmq, error) {
	c, err := LoadConfigSource()
	if err != nil {
		return nil, err
	}
	var rabbitmq Rabbitmq
	if err := c.Get("rabbitmq").Scan(&rabbitmq); err != nil {
		return nil, err
	}
	return &rabbitmq, nil
}
