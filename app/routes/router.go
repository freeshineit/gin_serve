package routes

import (
	"gin_serve/app/api"
	"gin_serve/app/middleware"
	"gin_serve/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) *gin.Engine {

	// 设置静态资源
	SetStaticFS(router)

	// set page router
	SetRoutesPage(router)

	apiGroup := router.Group("/api")

	{
		apiGroup.POST("/register", api.Register)
		apiGroup.POST("/login", api.Login)
		apiGroup.POST("/logout", api.Logout)

		apiGroup.GET("/query", func(c *gin.Context) {
			// message := c.Query("message")
			// nick := c.DefaultQuery("nick", "anonymous")
			c.JSON(http.StatusOK, models.BuildOKResponse(gin.H{
				"message": "message",
				"nick":    "nick",
			}))
		})

		// 注册需要带权限的路由
		RegisterGroup(apiGroup)
	}

	// /api/v1
	apiV1Group := router.Group("/api/v1", middleware.JwtAuth())
	{
		// register.go
		RegisterV1Group(apiV1Group)
	}

	// /api/v2
	apiV2Group := router.Group("/api/v2", middleware.JwtAuth())
	{
		// register.go
		RegisterV2Group(apiV2Group)
	}

	// socket /ws
	socketGroup := router.Group("/ws")
	{
		// register.go
		RegisterWsGroup(socketGroup)
	}

	return router
}
