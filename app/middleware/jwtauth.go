package middleware

import (
	"errors"
	"go_python_serve/app/models"
	"go_python_serve/app/utils"
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
			c.JSON(http.StatusUnauthorized, models.BuildErrorResponse("no token", errors.New("no token")))
			c.Abort()
			return
		}

		token = strings.Fields(token)[1]

		user, ok := utils.ParseToken(token)

		if !ok {
			c.JSON(http.StatusUnauthorized, models.BuildErrorResponse("token is expired", errors.New("token is expired")))
			c.Abort()
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
