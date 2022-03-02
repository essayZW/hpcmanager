package logic

import (
	"context"
	"fmt"

	"github.com/essayZW/hpcmanager/hpc/db"
	hpcdb "github.com/essayZW/hpcmanager/hpc/db"
	"github.com/essayZW/hpcmanager/hpc/source"
	"github.com/essayZW/hpcmanager/logger"
)

// HpcLogic hpc作业调度系统逻辑
type HpcLogic struct {
	hpcSource  source.HpcSource
	hpcUserDB  *hpcdb.HpcUserDB
	hpcGroupDB *hpcdb.HpcGroupDB
}

// AddUserWithGroup 创建组并添加用户到组
func (hpc *HpcLogic) AddUserWithGroup(ctx context.Context, username, groupname string) (map[string]interface{}, error) {
	res, err := hpc.hpcSource.AddUserWithGroup(username, groupname)
	if err != nil {
		return nil, err
	}
	success := res["success"]
	var status string
	status, ok := success.(string)
	if !ok {
		logger.Warn("AddUserWithGroup error: ", err)
		return nil, fmt.Errorf("AddUserWithGroup error with error %s", err.Error())
	}
	switch status {
	case "half":
		logger.Warn("AddUserWithGroup half success: ", res)
		fallthrough
	case "true":
		data, ok := res["data"]
		if !ok {
			logger.Warn("AddUserWithGroup error: ", err)
			return nil, fmt.Errorf("AddUserWithGroup error with error %s", err.Error())
		}
		mapData, ok := data.(map[string]interface{})
		if !ok {
			logger.Warn("AddUserWithGroup error: ", err)
			return nil, fmt.Errorf("AddUserWithGroup error with error %s", err.Error())
		}
		return mapData, nil
	default:
		logger.Warn("AddUserWithGroup error: ", err)
		return nil, fmt.Errorf("AddUserWithGroup error with error %s", err.Error())
	}
}

// AddUserToGroup 添加用户到现有的用户组中
func (hpc *HpcLogic) AddUserToGroup(ctx context.Context, username, groupname string, gid int) (map[string]interface{}, error) {
	res, err := hpc.hpcSource.AddUserToGroup(username, groupname, gid)
	if err != nil {
		return nil, err
	}
	success := res["success"]
	var status string
	status, ok := success.(string)
	if !ok {
		logger.Warn("AddUserToGroup error: ", err)
		return nil, fmt.Errorf("AddUserToGroup error with error %s", err.Error())
	}
	switch status {
	case "true":
		data, ok := res["data"]
		if !ok {
			logger.Warn("AddUserToGroup error: ", err)
			return nil, fmt.Errorf("AddUserToGroup error with error %s", err.Error())
		}
		mapData, ok := data.(map[string]interface{})
		if !ok {
			logger.Warn("AddUserToGroup error: ", err)
			return nil, fmt.Errorf("AddUserToGroup error with error %s", err.Error())
		}
		return mapData, nil
	default:
		logger.Warn("AddUserToGroup error: ", err)
		return nil, fmt.Errorf("AddUserToGroup error with error %s", err.Error())
	}
}

// CreateGroup 创建新的hpc节点上的用户组记录
func (hpc *HpcLogic) CreateGroup(ctx context.Context, groupName, queueName string, gid int) (int64, error) {
	return hpc.hpcGroupDB.Insert(ctx, &hpcdb.HpcGroup{
		Name:      groupName,
		GID:       gid,
		QueueName: queueName,
	})
}

// CreateUser 创建新的hpc节点上的用户记录
func (hpc *HpcLogic) CreateUser(ctx context.Context, userName string, uid int) (int64, error) {
	return hpc.hpcUserDB.Insert(ctx, &hpcdb.HpcUser{
		NodeUsername: userName,
		NodeUID:      uid,
	})
}

// NewHpc 创建一个HPC作业调度系统逻辑操作
func NewHpc(hpcSource source.HpcSource, hpcUserDB *db.HpcUserDB, hpcGroupDB *hpcdb.HpcGroupDB) *HpcLogic {
	return &HpcLogic{
		hpcSource:  hpcSource,
		hpcUserDB:  hpcUserDB,
		hpcGroupDB: hpcGroupDB,
	}
}
