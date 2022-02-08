package config

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/asim/go-micro/plugins/config/source/etcd/v4"
	"github.com/essayZW/hpcmanager"
	"github.com/essayZW/hpcmanager/logger"
	"go-micro.dev/v4/config"
)

// etcdDynamicConfig etcd为数据源的动态配置监听
type etcdDynamicConfig struct {
	conf config.Config
}

// Registry 注册
func (e *etcdDynamicConfig) Registry(path string, value interface{}, handler ValueChange) error {
	reflectValue := reflect.ValueOf(value)
	if reflectValue.Kind() != reflect.Ptr {
		return errors.New("value must be ptr")
	}

	if reflectValue.IsNil() {
		return errors.New("value can't be nil ptr")
	}

	watcher, err := e.conf.Watch(path)
	if err != nil {
		return err
	}
	go func() {
		for {
			newValue, err := watcher.Next()
			if err != nil {
				logger.Error("EtcdDynamicConfig watcher next error: ", err, " with path ", path)
			}
			logger.Info("EtcdDynamicConfig watcher path ", path, " changed")
			byt := newValue.Bytes()
			logger.Debug(string(byt))
			changedValue := make(map[string]interface{})
			err = json.Unmarshal(byt, &changedValue)
			if err != nil {
				logger.Error("EtcdDynamicConfig watcher next error: ", err, " with path ", path)
			}
			logger.Debug(changedValue)
			// 由于每次etcd上监听的数据发生变化时,都会有相应的Event事件
			// 修改值的时候会触发对应的一条事件
			// go-micro根据事件的数据更新本地存储的数据,该数据在Watch的时候初始化
			// 同时go-micro假定所有变化的数据值都是json格式的数据,因此变化的数据值必须是json格式,否则会解析错误
			// 因此假定监听的为 /path 则每次更改的是 /path/v 更改的数据格式为 { value: value },其中第二个value为实际更新的值
			v, ok := changedValue["v"]
			if !ok {
				logger.Debug("EtcdDynamicConfig changed value must be /path/v, but is ", path)
				continue
			}
			// 保证值必须是json格式
			changedValue, ok = v.(map[string]interface{})
			if !ok {
				logger.Debug("EtcdDynamicConfig changed value /path/v must be json data, but is ", changedValue)
				continue
			}
			if reallyValue, ok := changedValue["value"]; ok {
				// 防止数据类型不匹配导致崩溃从而程序退出以及监听停止
				go func() {
					defer func() {
						err := recover()
						if err != nil {
							logger.Error(err)
						}
					}()
					reflectValue.Elem().Set(reflect.ValueOf(reallyValue))
					if handler != nil {
						handler(reallyValue)
					}
				}()
				continue
			}
			logger.Debug("EtcdDynamicConfig changed value /path/v must have key value, but is ", changedValue)
		}
	}()
	return nil
}

// EtcdDynamicConfigPrefix etcd动态配置存储的前缀
const EtcdDynamicConfigPrefix = "/hpcmanager/micro/config"

// NewEtcd 创建一个基于etcd的动态配置工具
func NewEtcd() (DynamicConfig, error) {
	etcdSource := etcd.NewSource(
		etcd.WithAddress(hpcmanager.GetEtcdAddress()),
		etcd.WithPrefix(EtcdDynamicConfigPrefix),
		etcd.StripPrefix(true),
	)
	conf, err := config.NewConfig(
		config.WithSource(etcdSource),
	)
	if err != nil {
		logger.Error("New EtcdDynamicConfig error: ", err)
		return nil, err
	}
	return &etcdDynamicConfig{
		conf: conf,
	}, nil
}
