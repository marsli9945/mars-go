package marsGin

import (
	"github.com/go-playground/validator/v10"
	"github.com/marsli9945/mars-go/marsLog"
)

var validate *validator.Validate

func getValidate() *validator.Validate {
	if validate == nil {
		validate = validator.New()
	}
	return validate
}

// BindAndValid binds and validates data
func (g *Gin) BindAndValid(form any) {
	var err error
	err = g.Bind(form)
	if err != nil {
		marsLog.Logger().ErrorF("BindAndValid error: %v", err)
		panic(err)
	}
	err = getValidate().Struct(form)
	if err != nil {
		marsLog.Logger().ErrorF("BindAndValid error: %v", err)
		panic(err)
	}
}

func (g *Gin) BindQueryAndValid(form any) {
	var err error
	err = g.BindQuery(form)
	if err != nil {
		marsLog.Logger().ErrorF("BindQueryAndValid error: %v", err)
		panic(GetMsg(INVALID_PARAMS))
	}
	err = getValidate().Struct(form)
	if err != nil {
		marsLog.Logger().ErrorF("BindQueryAndValid error: %v", err)
		panic(err)
	}
}
