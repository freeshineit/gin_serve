package routes

import (
	"gin_serve/app/api"
	v1 "gin_serve/app/api/v1"
	v2 "gin_serve/app/api/v2"
	"gin_serve/app/api/ws"

	"github.com/gin-gonic/gin"
)

// register with auth api /api/xxx
func RegisterGroup(router *gin.RouterGroup) *gin.RouterGroup {

	router.POST("/file_upload", api.FileUpload)
	router.GET("/user/:id", api.GetUserByID)
	router.POST("/user", api.CreateUser)
	router.PUT("/user/:id", api.UpdateUser)
	router.DELETE("/user/:id", api.DeleteUser)

	return router
}

// register v1 api /api/v1/xxx
func RegisterV1Group(router *gin.RouterGroup) *gin.RouterGroup {

	// /api/v1/list
	router.GET("/list", v1.List)

	router.POST("/todo", v1.CreateTodo)       // /api/v1/todo [post]
	router.DELETE("/todo/:id", v1.DeleteTodo) // /api/v1/todo/:id [delete]
	router.PUT("/todo/:id", v1.PutTodo)       // /api/v1/todo/:id [put]
	router.GET("/todo/:id", v1.GetTodo)       // /api/v1/todo/:id [get]
	router.GET("/todos", v1.GetTodos)         // /api/v1/todos [get]

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
