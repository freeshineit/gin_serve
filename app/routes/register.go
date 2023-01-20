package routes

import (
	"gin_server/app/api"
	v1 "gin_server/app/api/v1"
	v2 "gin_server/app/api/v2"
	"gin_server/app/api/ws"

	"github.com/gin-gonic/gin"
)

// register with auth api /api/xxx
func RegisterGroup(router *gin.RouterGroup) *gin.RouterGroup {

	router.POST("/file_upload", api.FileUpload)

	return router
}

// register v1 api /api/v1/xxx
func RegisterV1Group(router *gin.RouterGroup) *gin.RouterGroup {

	// /api/v2/list
	router.GET("/list", v1.List)

	return router
}

// register v2 api /api/v2/xxx
func RegisterV2Group(router *gin.RouterGroup) *gin.RouterGroup {

	// /api/v2/list
	router.GET("/list", v2.List)
	return router
}

// register socket api /ws/xxx
func RegisterWsGroup(router *gin.RouterGroup) *gin.RouterGroup {
	router.GET("/ping", ws.Ping)
	return router
}
