package logic

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/fee/db"
	"github.com/essayZW/hpcmanager/logger"
)

// NodeDistributeBill 机器独占账单操作逻辑
type NodeDistributeBill struct {
	ndb *db.NodeDistributeBillDB

	mutex sync.Mutex
	// rate36CPU 36 核心节点费率/年
	rate36CPU float64
	// rate4GPU 4 gpu核心节点费率/年
	rate4GPU float64
	// rate8GPU 8 gpu核心节点费率/年
	rate8GPU float64
}

// Create 创建新的机器节点独占账单
func (ndbl *NodeDistributeBill) Create(
	ctx context.Context,
	applyID int,
	nodeDistributeID int,
	fee float64,
	userID int,
	username, userName string,
) (int64, error) {
	if applyID <= 0 {
		return 0, errors.New("invalid apply id")
	}
	if nodeDistributeID <= 0 {
		return 0, errors.New("invalid node distribute id")
	}
	if fee < 0 {
		return 0, errors.New("invalid fee")
	}
	if userID <= 0 {
		return 0, errors.New("invalid user id")
	}
	return ndbl.ndb.Insert(ctx, &db.NodeDistributeBill{
		ApplyID:          applyID,
		NodeDistributeID: nodeDistributeID,
		Fee:              fee,
		UserID:           userID,
		Username:         username,
		UserName:         userName,
		CreateTime:       time.Now(),
	})
}

// CalFee 计算账单的费用
func (ndbl *NodeDistributeBill) CalFee(startTimeUnix, endTimeUnix int64, nodeType string) (float64, error) {
	year := ndbl.calTimeDurationYear(startTimeUnix, endTimeUnix)
	var res float64
	var err error = nil
	switch nodeType {
	case "cpuc36":
		res = year * ndbl.rate36CPU
	case "gpuc4":
		res = year * ndbl.rate4GPU
	case "gpuc8":
		res = year * ndbl.rate8GPU
	default:
		res = 0
		err = errors.New("invalid nodeType")
	}
	return res, err
}

func (ndbl *NodeDistributeBill) calTimeDurationYear(startTimeUnix, endTimeUnix int64) float64 {
	startTime := time.Unix(startTimeUnix, 0)
	endTime := time.Unix(endTimeUnix, 0)
	// TODO: 对于费率的计算存疑惑,目前将时间差换算为年,然后乘对应的费率
	yearDuration := endTime.Year() - startTime.Year()
	var year float64
	year = float64(yearDuration)
	year += float64((endTime.Month()+12-startTime.Month())%12) / 12
	return year
}

// NewNodeDistributeBill 创建新的机器独占账单操作逻辑结构体
func NewNodeDistributeBill(ndb *db.NodeDistributeBillDB, dynamicConfig config.DynamicConfig) (*NodeDistributeBill, error) {
	res := &NodeDistributeBill{
		ndb: ndb,
	}

	// rate36CPU 36 核心节点费率/年
	var rate36CPU float64
	// rate4GPU 4 gpu核心节点费率/年
	var rate4GPU float64
	// rate8GPU 8 gpu核心节点费率/年
	var rate8GPU float64
	if err := dynamicConfig.Registry("fee_rate36CPU", &rate36CPU, func(newV interface{}) {
		res.mutex.Lock()
		defer res.mutex.Unlock()
		res.rate36CPU = rate36CPU
	}); err != nil {
		return nil, err
	}
	if err := dynamicConfig.Registry("fee_rate8GPU", &rate8GPU, func(newV interface{}) {
		res.mutex.Lock()
		defer res.mutex.Unlock()
		res.rate8GPU = rate8GPU
	}); err != nil {
		return nil, err
	}
	if err := dynamicConfig.Registry("fee_rate4GPU", &rate4GPU, func(newV interface{}) {
		res.mutex.Lock()
		defer res.mutex.Unlock()
		res.rate4GPU = rate4GPU
	}); err != nil {
		return nil, err
	}

	go func() {
		time.Sleep(time.Second)
		logger.Debug(res)
	}()
	return res, nil
}