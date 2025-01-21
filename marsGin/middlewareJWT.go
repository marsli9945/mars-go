package marsGin

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/marsli9945/mars-go/marsJwt"
)

// MiddlewareJWT is jwt middleware
func MiddlewareJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		appG := GetGin(c)

		token := c.GetHeader("Authorization")
		if token == "" {
			appG.Error(ERROR_AUTH)
			c.Abort()
			return
		}

		code = SUCCESS
		claims, err := marsJwt.ParseToken(token)
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				code = ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			default:
				code = ERROR_AUTH_CHECK_TOKEN_FAIL
			}
		} else {
			appG.SetUser(claims.Username)
		}

		if code != SUCCESS {
			appG.Error(code)
			c.Abort()
			return
		}

		c.Next()
	}
}
