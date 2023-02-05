package middleware

import (
	"gin_serve/helper"
	"gin_serve/message"
	"net/http"

	"github.com/gin-gonic/gin"
)

var TokenClaims = "TokenClaims"

// the jwt middleware
// JWTAuth jwt中间件
func JwtAuth() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// // cookie, _ := c.Cookie("authorization")
		// // if err != nil {
		// // 	c.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildErrorResponse(401, "no token", "no token found"))
		// // }
		// // log.Println("cookie", cookie)

		authorization := helper.GetHeaderToken(ctx)

		if authorization == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildErrorResponse(message.UnauthorizedCode, message.Unauthorized, message.Unauthorized))
			return
		}

		token, tokenClaims, err := helper.ValidateTokenAndClaims(authorization)

		if err == nil && token.Valid {
			ctx.Set(TokenClaims, tokenClaims)
			ctx.Next()
			return
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildErrorResponse(message.UnauthorizedCode, message.UnauthorizedExpired, message.UnauthorizedExpired))
		}
	}
}
