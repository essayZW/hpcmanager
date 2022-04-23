package json

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		createPaperAwardApplyParam := CreatePaperAwardApplyParam{}
		v.RegisterStructValidation(createPaperAwardApplyParam.Validator(), createPaperAwardApplyParam)
		checkPaperApplyParam := CheckPaperApplyParam{}
		v.RegisterStructValidation(checkPaperApplyParam.Validator(), checkPaperApplyParam)
		createTechnologyAwardApply := CreateTechnologyAwardApplyParam{}
		v.RegisterStructValidation(createTechnologyAwardApply.Validator(), createTechnologyAwardApply)
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

// CheckPaperApplyParam 审核论文奖励申请参数
type CheckPaperApplyParam struct {
	ID           int     `form:"id"           json:"id"           binding:"required"`
	CheckMoney   float64 `form:"checkMoney"   json:"checkMoney"`
	CheckMessage string  `form:"checkMessage" json:"checkMessage"`
	Accept       bool    `form:"accept"       json:"accept"`
}

func (param *CheckPaperApplyParam) Validator() validator.StructLevelFunc {
	return func(sl validator.StructLevel) {
		data := sl.Current().Interface().(CheckPaperApplyParam)
		if data.ID <= 0 {
			sl.ReportError(reflect.ValueOf(data.ID), "ID", "ID", "binding", "invalid apply id")
		}
		if data.CheckMoney < 0 {
			sl.ReportError(
				reflect.ValueOf(data.CheckMoney),
				"checkMoney",
				"checkMoney",
				"binding",
				"checkMoney can't less than 0",
			)
		}
	}
}

// CreateTechnologyAwardApplyParam 创建科技奖励申请参数
type CreateTechnologyAwardApplyParam struct {
	ProjectID      int    `form:"projectID"      json:"projectID"      binding:"required"`
	PrizeLevel     string `form:"prizeLevel"     json:"prizeLevel"     binding:"required"`
	PrizeImageName string `form:"prizeImageName" json:"prizeImageName" binding:"required"`
	RemarkMessage  string `form:"remarkMessage"  json:"remarkMessage"`
}

func (param *CreateTechnologyAwardApplyParam) Validator() validator.StructLevelFunc {
	return func(sl validator.StructLevel) {
		data := sl.Current().Interface().(CreateTechnologyAwardApplyParam)
		if data.ProjectID <= 0 {
			sl.ReportError(reflect.ValueOf(data.ProjectID), "ProjectID", "ProjectID", "binding", "invalid ProjectID")
		}
	}
}
