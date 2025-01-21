package marsGin

import (
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, msg string, data any) {
	g.JSON(httpCode, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
}

func (g *Gin) ResponseHttpOk(errCode int, msg string, data any) {
	g.Response(http.StatusOK, errCode, msg, data)
}

func (g *Gin) ResponseCode(errCode int, data any) {
	g.ResponseHttpOk(errCode, GetMsg(errCode), data)
}

func (g *Gin) Success(data any) {
	g.ResponseCode(SUCCESS, data)
}

func (g *Gin) Ok() {
	g.Success(nil)
}

func (g *Gin) Error(errCode int) {
	g.ResponseCode(errCode, nil)
}

func (g *Gin) ErrorMsg(msg string) {
	g.ResponseHttpOk(ERROR, msg, nil)
}
