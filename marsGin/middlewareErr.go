package marsGin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marsli9945/mars-go/marsLog"
	"runtime/debug"
)

func MiddlewareErr() gin.HandlerFunc {
	return TransformHandle(func(g *Gin) {
		defer func() {
			if err := recover(); err != nil {
				marsLog.Logger().Error(err, string(debug.Stack()))
				g.ErrorMsg(fmt.Sprintf("%s ", err))
				g.Context.Abort()
			}
		}()
		g.Context.Next()
	})
}
