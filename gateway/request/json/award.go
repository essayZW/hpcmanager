package json

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		createPaperAwardApplyParam := CreatePaperAwardApplyParam{}
		v.RegisterStructValidation(createPaperAwardApplyParam.Validator(), createPaperAwardApplyParam)
	}
}

// CreatePaperAwardApplyParam 创建论文奖励申请参数
type CreatePaperAwardApplyParam struct {
	Title               string `form:"title"               json:"title"               binding:"required"`
	Category            string `form:"category"            json:"category"            binding:"required"`
	Partition           string `form:"partition"           json:"partition"           binding:"required"`
	FirstPageImageName  string `form:"firstPageImageName"  json:"firstPageImageName"  binding:"required"`
	ThanksPageImageName string `form:"thanksPageImageName" json:"thanksPageImageName" binding:"required"`
	RemarkMessage       string `form:"remarkMessage"       json:"remarkMessage"`
}

func (param *CreatePaperAwardApplyParam) Validator() validator.StructLevelFunc {
	return func(sl validator.StructLevel) {
	}
}
