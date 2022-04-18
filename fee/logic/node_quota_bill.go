package logic

import (
	"context"
	"errors"
	"time"

	"github.com/essayZW/hpcmanager/config"
	"github.com/essayZW/hpcmanager/fee/db"
	"github.com/essayZW/hpcmanager/logger"
)

type NodeQuotaBill struct {
	nodeQuotaBillDB *db.NodeQuotaBillDB
}

// CreateNewBillParam 创建账单参数
type CreateNewBillParam struct {
	UserID      int
	UserName    string
	Username    string
	UserGroupID int
	OperType    QuoatOperationType
	OldSize     int
	NewSize     int
	OldEndTime  int64
	NewEndTime  int64
}

// CreateNewBill 创建新的账单
func (this *NodeQuotaBill) CreateNewBill(ctx context.Context, param *CreateNewBillParam) (int64, error) {
	if param == nil {
		return 0, errors.New("need param")
	}
	if param.UserID <= 0 {
		return 0, errors.New("invalid user id")
	}
	if param.UserGroupID <= 0 {
		return 0, errors.New("invalid user group id")
	}
	if param.OperType != ChangeQuotaSize && param.OperType != ChangeEndTime {
		return 0, errors.New("invalid oper type")
	}
	oldEndTime := time.Unix(param.OldEndTime, 0)
	newEndTime := time.Unix(param.NewEndTime, 0)

	fee := this.CalFee(param.OldSize, param.NewSize, oldEndTime, newEndTime)
	return this.nodeQuotaBillDB.Insert(ctx, &db.NodeQuotaBill{
		UserID:      param.UserID,
		Username:    param.Username,
		UserName:    param.UserName,
		UserGroupID: param.UserGroupID,
		OperType:    int8(param.OperType),
		OldSize:     param.OldSize,
		NewSize:     param.NewSize,
		OldEndTime:  oldEndTime,
		NewEndTime:  newEndTime,
		Fee:         fee,
		CreateTime:  time.Now(),
	})
}

// CalFee 计算费用
func (this *NodeQuotaBill) CalFee(oldSize, newSize int, oldEndTime, newEndTime time.Time) float64 {
	// TODO: need implement
	logger.Fatal("need implement")
	return 0
}

// NewNodeQuotaBill 创建新的节点存储账单操作逻辑
func NewNodeQuotaBill(nodeQuotaBillDB *db.NodeQuotaBillDB, dynamicConfig config.DynamicConfig) (*NodeQuotaBill, error) {
	return &NodeQuotaBill{
		nodeQuotaBillDB: nodeQuotaBillDB,
	}, nil
}

// QuotaOperationType 存储操作类型
type QuoatOperationType int8

const (
	// ChangeQuotaSize 修改容量
	ChangeQuotaSize QuoatOperationType = 1
	// ChangeEndTime 修改期限的最后的时间
	ChangeEndTime QuoatOperationType = 2
)
