package middleware

import (
	"gin_serve/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// the jwt middleware
// JWTAuth jwt中间件
func JwtAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		// cookie, _ := c.Cookie("authorization")
		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildErrorResponse(401, "no token", "no token found"))
		// }
		// log.Println("cookie", cookie)

		if len(authorization) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildErrorResponse(401, "no token", "no token found"))
			return
		}

		authorization = strings.Fields(authorization)[1]
		// log.Println("authorization", authorization)

		token, err := helper.ValidateToken(authorization)

		if err == nil && token.Valid {
			// c.Set("user", user)
			c.Next()
			return
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildErrorResponse(401, "token is expired", "token is expired"))
		}

	}
}
