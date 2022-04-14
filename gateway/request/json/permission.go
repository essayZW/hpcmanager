package json

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		createCommonAdminParam := &ChangeUserPermissionParam{}
		v.RegisterStructValidation(createCommonAdminParam.Validator(), &createCommonAdminParam)
	}
}

// ChangeUserPermissionParam 修改普通管理员用户请求参数
type ChangeUserPermissionParam struct {
	UserID int `form:"userID" json:"userID" binding:"required"`
}

// Validator 验证
func (param *ChangeUserPermissionParam) Validator() validator.StructLevelFunc {
	return func(sl validator.StructLevel) {
		data := sl.Current().Interface().(ChangeUserPermissionParam)
		if data.UserID <= 0 {
			sl.ReportError(
				reflect.ValueOf(data.UserID),
				"userID",
				"userID",
				"binding",
				"invalid userID",
			)
		}
	}
}
