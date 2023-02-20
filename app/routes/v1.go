package routes

import (
	v1 "gin_serve/app/api/v1"
	"gin_serve/app/middleware"

	"github.com/gin-gonic/gin"
)

// register v1 api /api/v1/xxx
func RegisterV1Group(router *gin.RouterGroup) *gin.RouterGroup {
	return router
}

// register v1 api /api/v1/xxx with auth
func RegisterV1GroupWithAuth(router *gin.RouterGroup) *gin.RouterGroup {

	authRouter := router.Group("/", middleware.JwtAuth())

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
