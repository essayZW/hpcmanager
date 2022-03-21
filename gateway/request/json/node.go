package json

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		createNodeApplyParam := CreateNodeApplyParam{}
		v.RegisterStructValidation(createNodeApplyParam.Validator(), &createNodeApplyParam)
		checkNodeApplyParam := CheckNodeApplyParam{}
		v.RegisterStructValidation(checkNodeApplyParam.Validator(), &checkNodeApplyParam)
	}
}

// CreateNodeApplyParam 创建机器节点申请记录请求参数
type CreateNodeApplyParam struct {
	ProjectID int    `form:"projectID" json:"projectID" binding:"required"`
	NodeType  string `form:"nodeType" json:"nodeType" binding:"required"`
	NodeNum   int    `form:"nodeNum" json:"nodeNum" binding:"required"`
	StartTime int64  `form:"startTime" json:"startTime" binding:"required"`
	EndTime   int64  `form:"endTime" json:"endTime" binding:"required"`
}

// Validator 验证器
func (c *CreateNodeApplyParam) Validator() validator.StructLevelFunc {
	return func(sl validator.StructLevel) {
		data := sl.Current().Interface().(CreateNodeApplyParam)

		if data.ProjectID <= 0 {
			sl.ReportError(reflect.ValueOf(data.ProjectID), "projectID", "projectID", "binding", "projectID invalid")
		}
		if data.NodeNum <= 0 {
			sl.ReportError(reflect.ValueOf(data.NodeNum), "nodeNum", "nodeNum", "binding", "nodeNum invalid")
		}
	}
}

// CheckNodeApplyParam 审核机器节点申请信息
type CheckNodeApplyParam struct {
	ApplyID int `form:"applyID" json:"applyID" binding:"required"`
	// CheckStatus 默认不存在则为false
	CheckStatus  bool   `form:"checkStatus" json:"checkStatus"`
	CheckMessage string `form:"checkMessage" json:"checkMessage"`
	// TutorCheck 默认不存在则为false
	TutorCheck bool `form:"tutorCheck" json:"tutorCheck"`
}

// Validator 验证器
func (c *CheckNodeApplyParam) Validator() validator.StructLevelFunc {
	return func(sl validator.StructLevel) {
		data := sl.Current().Interface().(CheckNodeApplyParam)

		if data.ApplyID <= 0 {
			sl.ReportError(reflect.ValueOf(data.ApplyID), "applyID", "applyID", "binding", "applyID error")
		}
	}
}
