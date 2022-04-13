package json

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		payNodeDistributeBill := PayNodeDistributeBillParam{}
		v.RegisterStructValidation(payNodeDistributeBill.Validator(), &payNodeDistributeBill)
	}
}

// PayNodeDistributeBillParam 支付机器独占账单参数
type PayNodeDistributeBillParam struct {
	ID         int     `form:"id"         json:"id"         binding:"required"`
	PayMoney   float64 `form:"payMoney"   json:"payMoney"   binding:"required"`
	PayType    float64 `form:"payType"    json:"payType"    binding:"required"`
	PayMessage string  `form:"payMessage" json:"payMessage" binding:"required"`
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
