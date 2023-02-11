package routes

import (
	"gin_serve/app/api"
	socket "gin_serve/app/api/socket"
	v1 "gin_serve/app/api/v1"
	v2 "gin_serve/app/api/v2"
	"gin_serve/app/middleware"
	"gin_serve/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

var jwtAuthMiddleware = middleware.JwtAuth()

// register with auth api /api/xxx
func RegisterGroup(router *gin.RouterGroup) *gin.RouterGroup {
	router.POST("/register", api.Register)
	router.POST("/login", api.Login)
	router.POST("/logout", api.Logout)

	router.GET("/query", func(c *gin.Context) {
		// message := c.Query("message")
		// nick := c.DefaultQuery("nick", "anonymous")
		c.JSON(http.StatusOK, helper.BuildResponse("success", gin.H{
			"message": "message",
			"nick":    "nick",
		}))
	})

	return router
}

// register with auth api /api/xxx
func RegisterGroupWithAuth(router *gin.RouterGroup) *gin.RouterGroup {
	authRouter := router.Group("/", jwtAuthMiddleware)

	authRouter.POST("/file_upload", api.FileUpload) // /api/file_upload [post]
	authRouter.GET("/user/:id", api.GetUserByID)    // /api/user/:id [get]
	authRouter.POST("/user", api.CreateUser)        // /api/user [post]
	authRouter.PUT("/user/:id", api.UpdateUser)     // /api/user/:id [put]
	authRouter.DELETE("/user/:id", api.DeleteUser)  // /api/user/:id [delete]

	authRouter.GET("/auth", api.AuthPage) // /api/auth [get]

	return authRouter
}

// register v1 api /api/v1/xxx
func RegisterV1Group(router *gin.RouterGroup) *gin.RouterGroup {
	return router
}

// register v1 api /api/v1/xxx with auth
func RegisterV1GroupWithAuth(router *gin.RouterGroup) *gin.RouterGroup {

	authRouter := router.Group("/", jwtAuthMiddleware)

	// /api/v1/list
	authRouter.GET("/list", v1.List) // /api/v1/list

	// todo
	authRouter.POST("/todo", v1.CreateTodo)                // /api/v1/todo [post]
	authRouter.DELETE("/todo/:id", v1.DeleteTodo)          // /api/v1/todo/:id [delete]
	authRouter.PUT("/todo/:id/content", v1.PutTodoContent) // /api/v1/todo/:id/content [put]
	authRouter.PUT("/todo/:id/status", v1.PutTodoStatus)   // /api/v1/todo/:id/status [put]
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
	authRouter := router.Group("/", jwtAuthMiddleware)

	authRouter.GET("/list", v2.List) // /api/v2/list
	return router
}

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
