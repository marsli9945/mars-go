package marsGin

import (
	"github.com/gin-gonic/gin"
)

// MiddlewareCors 跨域设置
func MiddlewareCors() gin.HandlerFunc {
	return TransformHandle(func(g *Gin) {
		method := g.Context.Request.Method
		origin := g.Context.Request.Header.Get("Origin") // 请求头部
		if origin != "" {
			g.Context.Header("Access-Control-Allow-Origin", origin) // 可将将 * 替换为指定的域名
			g.Context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			g.Context.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			g.Context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			g.Context.Header("Access-Control-Allow-Credentials", "true")
		}
		// 允许类型校验
		if method == "OPTIONS" {
			g.Context.AbortWithStatus(204)
			return
		}
		g.Context.Next()
	})
}
