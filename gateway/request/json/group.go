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
		createJoinGroupApplyParam := CreateJoinGroupApplyParam{}
		v.RegisterStructValidation(createJoinGroupApplyParam.Validator(), &createJoinGroupApplyParam)
		checkJoinGroupApplyParam := CheckJoinGroupApplyParam{}
		v.RegisterStructValidation(checkJoinGroupApplyParam.Validator(), &checkJoinGroupApplyParam)
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

// CreateJoinGroupApplyParam 创建加入组申请参数
type CreateJoinGroupApplyParam struct {
	ApplyGroupID int `form:"applyGroupID" json:"applyGroupID" binding:"required"`
}

// Validator 验证器
func (c *CreateJoinGroupApplyParam) Validator() validator.StructLevelFunc {
	return func(sl validator.StructLevel) {
		data := sl.Current().Interface().(CreateJoinGroupApplyParam)
		if data.ApplyGroupID <= 0 {
			sl.ReportError(reflect.ValueOf(data.ApplyGroupID), "applyGroupID", "applyGroupID", "binding", "applyGroupID error")
		}
	}
}

// CheckJoinGroupApplyParam 审核
type CheckJoinGroupApplyParam struct {
	ApplyID int `form:"applyID" json:"applyID" binding:"required"`
	// CheckStatus 默认不存在则为false
	CheckStatus  bool   `form:"checkStatus" json:"checkStatus"`
	CheckMessage string `form:"checkMessage" json:"checkMessage"`
	// TutorCheck 默认不存在则为false
	TutorCheck bool `form:"tutorCheck" json:"tutorCheck"`
}

// Validator 验证器
func (c *CheckJoinGroupApplyParam) Validator() validator.StructLevelFunc {
	return func(sl validator.StructLevel) {
		data := sl.Current().Interface().(CheckJoinGroupApplyParam)

		if data.ApplyID <= 0 {
			sl.ReportError(reflect.ValueOf(data.ApplyID), "applyID", "applyID", "binding", "applyID error")
		}
	}
}
