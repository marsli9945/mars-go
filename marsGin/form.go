package marsGin

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func getValidate() *validator.Validate {
	if validate == nil {
		validate = validator.New()
	}
	return validate
}

// BindAndValid binds and validates data
func (g *Gin) BindAndValid(form interface{}) {
	var err error
	err = g.C.Bind(form)
	if err != nil {
		panic(err)
	}
	err = getValidate().Struct(form)
	if err != nil {
		panic(err)
	}
}

func (g *Gin) BindQueryAndValid(form interface{}) {
	var err error
	err = g.C.BindQuery(form)
	if err != nil {
		panic(GetMsg(INVALID_PARAMS))
	}
	err = getValidate().Struct(form)
	if err != nil {
		panic(err)
	}
}
