package json

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		createCommonAdminParam := &CreateCommonAdminParam{}
		v.RegisterStructValidation(createCommonAdminParam.Validator(), &createCommonAdminParam)
	}
}

// CreateCommonAdminParam 创建新的普通管理员用户请求参数
type CreateCommonAdminParam struct {
	UserID int `form:"userID" json:"userID" binding:"required"`
}

// Validator 验证
func (param *CreateCommonAdminParam) Validator() validator.StructLevelFunc {
	return func(sl validator.StructLevel) {
		data := sl.Current().Interface().(CreateCommonAdminParam)
		if data.UserID <= 0 {
			sl.ReportError(reflect.ValueOf(data.UserID), "userID", "userID", "binding", "invalid userID")
		}
	}
}
