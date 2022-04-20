package json

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		payNodeDistributeBill := PayNodeDistributeBillParam{}
		v.RegisterStructValidation(payNodeDistributeBill.Validator(), payNodeDistributeBill)
		payGroupNodeUsageBill := PayGroupNodeUsageBillParam{}
		v.RegisterStructValidation(payGroupNodeUsageBill.Validator(), payGroupNodeUsageBill)
	}
}

// PayNodeDistributeBillParam 支付机器独占账单参数
type PayNodeDistributeBillParam struct {
	ID         int     `form:"id"         json:"id"         binding:"required"`
	PayMoney   float64 `form:"payMoney"   json:"payMoney"`
	PayType    float64 `form:"payType"    json:"payType"    binding:"required"`
	PayMessage string  `form:"payMessage" json:"payMessage"`
}

func (param *PayNodeDistributeBillParam) Validator() validator.StructLevelFunc {
	return func(sl validator.StructLevel) {
		data := sl.Current().Interface().(PayNodeDistributeBillParam)
		if data.PayMoney < 0 {
			sl.ReportError(reflect.ValueOf(data.PayMoney), "payType", "payType", "binding", "invalid payType")
		}
		if data.PayType != 1 && data.PayType != 2 {
			sl.ReportError(reflect.ValueOf(data.PayType), "payType", "payType", "binding", "invalid payType")
		}
	}
}

// PayGroupNodeUsageBillParam 支付用户组机器节点时长账单参数
type PayGroupNodeUsageBillParam struct {
	UserGroupID int     `form:"userGroupID" json:"userGroupID" binding:"required"`
	PayType     int     `form:"payType"     json:"payType"     binding:"required"`
	PayMessage  string  `form:"payMessage"  json:"payMessage"`
	NeedFee     float64 `form:"needFee"     json:"needFee"     binding:"required"`
}

func (param *PayGroupNodeUsageBillParam) Validator() validator.StructLevelFunc {
	return func(sl validator.StructLevel) {
		data := sl.Current().Interface().(PayGroupNodeUsageBillParam)
		if data.PayType != 1 && data.PayType != 2 {
			sl.ReportError(reflect.ValueOf(data.PayType), "payType", "payType", "binding", "invalid payType")
		}
		if data.NeedFee < 0 {
			sl.ReportError(reflect.ValueOf(data.NeedFee), "needFee", "needFee", "binding", "invalid needFee")
		}
	}
}

// PayNodeQuotaBillParam 支付机器存储账单参数
type PayNodeQuotaBillParam struct {
	BillID     int     `form:"billID"     json:"billID"     binding:"required"`
	PayType    int     `form:"payType"    json:"payType"    binding:"required"`
	PayMessage string  `form:"payMessage" json:"payMessage"`
	PayMoney   float64 `form:"payMoney"   json:"payMoney"`
}

func (param *PayNodeQuotaBillParam) Validator() validator.StructLevelFunc {
	return func(sl validator.StructLevel) {
		data := sl.Current().Interface().(PayNodeQuotaBillParam)

		if data.BillID <= 0 {
			sl.ReportError(reflect.ValueOf(data.BillID), "billID", "billID", "binding", "invalid billID")
		}
		if data.PayType != 1 && data.PayType != 2 {
			sl.ReportError(reflect.ValueOf(data.PayType), "payType", "payType", "binding", "invalid payType")
		}
	}
}
