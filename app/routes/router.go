package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) *gin.Engine {

	// 设置静态资源
	SetStaticFS(router)

	// set page router
	SetRoutesPage(router)

	apiGroup := router.Group("/api")
	{
		RegisterGroup(apiGroup)
		// 注册需要带权限的路由
		RegisterGroupWithAuth(apiGroup)
	}

	// /api/v1
	apiV1Group := router.Group("/api/v1")
	{
		// register.go
		RegisterV1Group(apiV1Group)
		RegisterV1GroupWithAuth(apiV1Group)
	}

	// /api/v2
	apiV2Group := router.Group("/api/v2")
	{
		// register.go
		RegisterV2Group(apiV2Group)
		RegisterV2GroupWithAuth(apiV2Group)
	}
	return router
}

func SetupSocketRoutes(router *gin.Engine) *gin.Engine {
	// socket /ws
	socketGroup := router.Group("/ws")
	{
		// register.go
		RegisterWsGroup(socketGroup)
		RegisterV2GroupWithAuth(socketGroup)
	}

	return router
}
