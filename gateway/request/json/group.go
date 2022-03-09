package json

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		createUserParam := CreateGroupParam{}
		v.RegisterStructValidation(createUserParam.Validator(), &createUserParam)
	}
}

// CreateGroupParam 创建组的参数
type CreateGroupParam struct {
	GroupName string `form:"groupName" json:"groupName" binding:"required"`
	QueueName string `form:"queueName" json:"queueName" binding:"required"`
	TutorID   int    `form:"tutorID" json:"tutorID" binding:"required"`
}

// Validator 验证器
func (cgp *CreateGroupParam) Validator() validator.StructLevelFunc {
	return func(sl validator.StructLevel) {
		data := sl.Current().Interface().(CreateGroupParam)
		if data.TutorID == 0 {
			sl.ReportError(reflect.ValueOf(data.TutorID), "tutorID", "tutorID", "binding", "tutorID error")
		}
		if len(data.GroupName) == 0 || len(data.GroupName) > 64 {
			sl.ReportError(reflect.ValueOf(data.GroupName), "groupName", "groupName", "binding", "groupName length error")
		}
		if len(data.QueueName) == 0 || len(data.QueueName) > 64 {
			sl.ReportError(reflect.ValueOf(data.QueueName), "queueName", "queueName", "binding", "queueName length error")
		}
	}
}
