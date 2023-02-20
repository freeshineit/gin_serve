package routes

import (
	v2 "gin_serve/app/api/v2"
	"gin_serve/app/middleware"

	"github.com/gin-gonic/gin"
)

// register v2 api /api/v2/xxx
func RegisterV2Group(router *gin.RouterGroup) *gin.RouterGroup {
	return router
}

// register v2 api /api/v2/xxx with auth
func RegisterV2GroupWithAuth(router *gin.RouterGroup) *gin.RouterGroup {
	authRouter := router.Group("/", middleware.JwtAuth())

	authRouter.GET("/list", v2.List) // /api/v2/list
	return router
}
