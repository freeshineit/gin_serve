package routes

import (
	socket "gin_serve/app/api/socket"

	"github.com/gin-gonic/gin"
)

// register socket api /ws/xxx
func RegisterWsGroup(router *gin.RouterGroup) *gin.RouterGroup {
	router.GET("/ping", socket.Ping) // /ws/ping
	return router
}

// register socket api /ws/xxx with auth
func RegisterWsGroupWithAuth(router *gin.RouterGroup) *gin.RouterGroup {
	// router.Group("/", jwtAuthMiddleware)
	return router
}
