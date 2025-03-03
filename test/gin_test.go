package test

import (
	"github.com/gin-gonic/gin"
	"github.com/marsli9945/mars-go/marsGin"
	"net/http"
	"testing"
)

func hello(g *marsGin.Gin) {
	g.Ok()
}

func TestTransform(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), marsGin.MiddlewareErr(), marsGin.MiddlewareCors())

	// 示例路由
	r.GET("/custom", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Custom endpoint"})
	})

	r.GET("/error", func(c *gin.Context) {
		panic("test error")
	})

	r.GET("/hello", marsGin.TransformHandle(hello))

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
