package routes

import (
	"go_python_serve/app/api"
	"go_python_serve/app/middleware"
	"go_python_serve/app/models"
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
		RegisterWithAuthGroup(apiGroup.Group("", middleware.JwtAuth()))
	}

	// /api/v1
	apiV1Group := router.Group("/api/v1", middleware.JwtAuth())
	{

		RegisterV1Group(RegisterWithAuthGroup(apiV1Group))
	}

	// /api/v2
	apiV2Group := router.Group("/api/v2", middleware.JwtAuth())
	{
		RegisterV2Group(RegisterWithAuthGroup(apiV2Group))
	}

	return router
}
