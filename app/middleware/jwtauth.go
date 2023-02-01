package middleware

import (
	"gin_serve/app/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// the jwt middleware
// JWTAuth jwt中间件
func JwtAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if len(token) == 0 {
			c.JSON(http.StatusUnauthorized, utils.BuildErrorResponse(1, "no token", "no token"))
			c.Abort()
			return
		}

		token = strings.Fields(token)[1]

		user, ok := utils.ParseToken(token)

		if !ok {
			c.JSON(http.StatusUnauthorized, utils.BuildErrorResponse(1, "token is expired", "token is expired"))
			c.Abort()
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
