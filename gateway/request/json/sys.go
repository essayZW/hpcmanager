package json

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		setCasConfigParam := SetCasConfigParam{}
		v.RegisterStructValidation(setCasConfigParam.Validator(), setCasConfigParam)
	}
}

type SetCasConfigParam struct {
	Enable     bool   `form:"enable"     json:"enable"`
	AuthServer string `form:"authServer" json:"authServer"`
}

func (param *SetCasConfigParam) Validator() validator.StructLevelFunc {
	return func(sl validator.StructLevel) {

	}
}
