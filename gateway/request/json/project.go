package json

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		createProjectParam := CreateProjectParam{}
		v.RegisterStructValidation(createProjectParam.Validator(), &createProjectParam)
	}
}

// CreateProjectParam 创建新的项目记录请求参数
type CreateProjectParam struct {
	Name        string `form:"name"        json:"name"        binding:"required"`
	From        string `form:"from"        json:"from"`
	Numbering   string `form:"numbering"   json:"numbering"`
	Expenses    string `form:"expenses"    json:"expenses"`
	Description string `form:"description" json:"description"`
}

// Validator 验证器
func (cpp *CreateProjectParam) Validator() validator.StructLevelFunc {
	return func(sl validator.StructLevel) {
		data := sl.Current().Interface().(CreateProjectParam)
		if data.Name == "" {
			sl.ReportError(
				reflect.ValueOf(data.Name),
				"name",
				"name",
				"binding",
				"project name error",
			)
		}
	}
}
