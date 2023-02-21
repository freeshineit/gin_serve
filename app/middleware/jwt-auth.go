package middleware

import (
	"gin_serve/app/service"
	"gin_serve/config"
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

		authorization := helper.GetHeaderToken(ctx)

		if authorization == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildErrorResponse(message.UnauthorizedCode, message.Unauthorized, message.Unauthorized))
			return
		}

		jwtService := service.NewJWTService(config.RedisClient)

		if jwtService.IsInBlacklist(authorization) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildErrorResponse(message.UnauthorizedCode, message.UnauthorizedExpired, message.UnauthorizedExpired))
			return
		}

		tokenClaims, valid, err := helper.ValidateTokenAndBackClaims(authorization)

		if err == nil && valid {
			ctx.Set(TokenClaims, tokenClaims)
			ctx.Next()
			return
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.BuildErrorResponse(message.UnauthorizedCode, message.UnauthorizedExpired, message.UnauthorizedExpired))
			return
		}
	}
}
