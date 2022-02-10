package json

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		login := Login{}
		v.RegisterStructValidation(login.Validator(), login)
	}
}

// Login 包含了用户名和密码的请求参数结构
type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// Validator 验证
func (login *Login) Validator() validator.StructLevelFunc {
	return func(sl validator.StructLevel) {
		data := sl.Current().Interface().(Login)
		if len(data.Username) < 6 || len(data.Username) > 32 {
			sl.ReportError(reflect.ValueOf(data.Username), "username", "username", "binding", "username length error")
		}
		if len(data.Password) <= 0 || len(data.Password) > 16 {
			sl.ReportError(reflect.ValueOf(data.Password), "password", "password", "binding", "username length error")
		}
	}
}
