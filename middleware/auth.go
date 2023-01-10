package middleware

import (
	"go_python_serve/utils"

	"github.com/gin-gonic/gin"
)

// the jwt middleware
func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		utils.ParseToken(auth)
	}
}
