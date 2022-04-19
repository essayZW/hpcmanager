package json

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		setUserQuotaParam := &SetUserQuotaParam{}
		v.RegisterStructValidation(setUserQuotaParam.Validator(), &setUserQuotaParam)
	}
}

// SetUserQuotaParam 设置用户存储信息参数
type SetUserQuotaParam struct {
	HpcUserID           int   `form:"hpcUserID"           json:"hpcUserID"           binding:"required"`
	OldSize             int   `form:"oldSize"             json:"oldSize"             binding:"required"`
	NewSize             int   `form:"newSize"             json:"newSize"             binding:"required"`
	OldEndTimeMilliUnix int64 `form:"oldEndTimeMilliUnix" json:"oldEndTimeMilliUnix" binding:"required"`
	NewEndTimeMilliUnix int64 `form:"newEndTimeMilliUnix" json:"newEndTimeMilliUnix" binding:"required"`
	ModifyData          bool  `form:"modifyData"          json:"modifyData"`
}

func (param *SetUserQuotaParam) Validator() validator.StructLevelFunc {
	return func(sl validator.StructLevel) {
		data := sl.Current().Interface().(SetUserQuotaParam)
		if data.HpcUserID <= 0 {
			sl.ReportError(reflect.ValueOf(data.HpcUserID), "hpcUserID", "hpcUserID", "binding", "invalid hpcUserID")
		}
		if data.NewSize < data.OldSize {
			sl.ReportError(reflect.ValueOf(data.NewSize), "NewSize", "NewSize", "binding", "不能减少最大的容量")
		}
		if data.NewEndTimeMilliUnix < data.OldEndTimeMilliUnix {
			sl.ReportError(
				reflect.ValueOf(data.NewEndTimeMilliUnix),
				"NewEndTimeMilliUnix",
				"NewEndTimeMilliUnix",
				"binding",
				"不能减少期限",
			)
		}
	}
}
