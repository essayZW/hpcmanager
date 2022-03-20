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
