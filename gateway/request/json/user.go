package json

import (
	"reflect"
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		login := Login{}
		v.RegisterStructValidation(login.Validator(), login)

		createUserParam := CreateUserParam{}
		v.RegisterStructValidation(createUserParam.Validator(), createUserParam)
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

// CreateUserParam 创建用户参数
type CreateUserParam struct {
	Login
	Tel         string `form:"tel" json:"tel"`
	Email       string `form:"email" json:"email"`
	Name        string `form:"name" json:"name" binding:"required"`
	CollegeName string `form:"collegeName" json:"collegeName"`
}

// Validator 参数验证
func (user *CreateUserParam) Validator() validator.StructLevelFunc {
	return func(sl validator.StructLevel) {
		data := sl.Current().Interface().(CreateUserParam)
		if len(data.Username) < 6 || len(data.Username) > 32 {
			sl.ReportError(reflect.ValueOf(data.Username), "username", "username", "binding", "username length error")
		}
		if len(data.Password) <= 0 || len(data.Password) > 16 {
			sl.ReportError(reflect.ValueOf(data.Password), "password", "password", "binding", "username length error")
		}
		pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
		reg := regexp.MustCompile(pattern)
		if data.Email != "" && !reg.MatchString(data.Email) {
			sl.ReportError(reflect.ValueOf(data.Email), "email", "email", "binding", "invalid email")
		}
	}
}

// UpdateUserInfoParam 用户信息更新参数
type UpdateUserInfoParam struct {
	ID      int    `form:"id" json:"id"`
	Tel     string `form:"tel" json:"tel"`
	Email   string `form:"email" json:"email"`
	College string `form:"college" json:"college"`
}
