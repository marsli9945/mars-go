package marsGin

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	*gin.Context
}

func TransformHandle(handle func(g *Gin)) func(c *gin.Context) {
	return func(c *gin.Context) {
		handle(&Gin{c})
	}
}
