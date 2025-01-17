package marsGin

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

func GetGin(c *gin.Context) *Gin {
	return &Gin{
		C: c,
	}
}
