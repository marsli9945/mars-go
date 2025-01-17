package marsGin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marsli9945/mars-go/marsLog"
	"runtime/debug"
)

func MiddlewareErr() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			appG := GetGin(c)
			if err := recover(); err != nil {
				marsLog.Logger().Error(err, string(debug.Stack()))
				appG.ErrorMsg(fmt.Sprintf("%s ", err))
				c.Abort()
			}
		}()
		c.Next()
	}
}
