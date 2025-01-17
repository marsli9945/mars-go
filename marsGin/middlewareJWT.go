package marsGin

import (
	"github.com/dgrijalva/jwt-go"
	util "github.com/marsli9945/mars-go/marsJwt"

	"github.com/gin-gonic/gin"
)

// MiddlewareJWT is jwt middleware
func MiddlewareJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		appG := Gin{C: c}

		code = SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			code = ERROR_AUTH
		} else {
			claims, err := util.ParseToken(token)
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
		}

		if code != SUCCESS {
			appG.Error(code)

			c.Abort()
			return
		}

		c.Next()
	}
}
