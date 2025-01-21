package marsGin

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	*gin.Context
}

func GetGin(c *gin.Context) *Gin {
	return &Gin{c}
}
