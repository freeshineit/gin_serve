package routes

import (
	"gin_serve/app/api"
	v1 "gin_serve/app/api/v1"
	v2 "gin_serve/app/api/v2"
	"gin_serve/app/api/ws"
	"gin_serve/app/middleware"
	"gin_serve/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// register with auth api /api/xxx
func RegisterGroup(router *gin.RouterGroup) *gin.RouterGroup {
	router.POST("/register", api.Register)
	router.POST("/login", api.Login)
	router.POST("/logout", api.Logout)

	router.GET("/query", func(c *gin.Context) {
		// message := c.Query("message")
		// nick := c.DefaultQuery("nick", "anonymous")
		c.JSON(http.StatusOK, models.BuildOKResponse(gin.H{
			"message": "message",
			"nick":    "nick",
		}))
	})

	return router
}

// register with auth api /api/xxx
func RegisterGroupWithAuth(router *gin.RouterGroup) *gin.RouterGroup {
	authRouter := router.Group("/", middleware.JwtAuth())

	authRouter.POST("/file_upload", api.FileUpload)
	authRouter.GET("/user/:id", api.GetUserByID)
	authRouter.POST("/user", api.CreateUser)
	authRouter.PUT("/user/:id", api.UpdateUser)
	authRouter.DELETE("/user/:id", api.DeleteUser)

	authRouter.GET("/auth", api.AuthPage)

	return authRouter
}

// register v1 api /api/v1/xxx
func RegisterV1Group(router *gin.RouterGroup) *gin.RouterGroup {
	return router
}

// register v1 api /api/v1/xxx with auth
func RegisterV1GroupWithAuth(router *gin.RouterGroup) *gin.RouterGroup {

	authRouter := router.Group("/", middleware.JwtAuth())

	// /api/v1/list
	authRouter.GET("/list", v1.List)

	authRouter.POST("/todo", v1.CreateTodo)                // /api/v1/todo [post]
	authRouter.DELETE("/todo/:id", v1.DeleteTodo)          // /api/v1/todo/:id [delete]
	authRouter.PUT("/todo/:id/content", v1.PutTodoContent) // /api/v1/todo/:id [put]
	authRouter.PUT("/todo/:id/status", v1.PutTodoStatus)   // /api/v1/todo/:id [put]
	authRouter.GET("/todo/:id", v1.GetTodo)                // /api/v1/todo/:id [get]
	authRouter.GET("/todos", v1.GetTodos)                  // /api/v1/todos [get]

	return authRouter
}

// register v2 api /api/v2/xxx
func RegisterV2Group(router *gin.RouterGroup) *gin.RouterGroup {
	return router
}

// register v2 api /api/v2/xxx with auth
func RegisterV2GroupWithAuth(router *gin.RouterGroup) *gin.RouterGroup {
	authRouter := router.Group("/", middleware.JwtAuth())

	// /api/v2/list
	authRouter.GET("/list", v2.List)
	return router
}

// register socket api /ws/xxx
func RegisterWsGroup(router *gin.RouterGroup) *gin.RouterGroup {
	router.GET("/ping", ws.Ping)
	return router
}

// register socket api /ws/xxx with auth
func RegisterWsGroupWithAuth(router *gin.RouterGroup) *gin.RouterGroup {
	router.Group("/", middleware.JwtAuth())

	return router
}
