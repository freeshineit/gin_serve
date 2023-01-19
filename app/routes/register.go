package routes

import (
	"go_python_serve/app/api"
	v1 "go_python_serve/app/api/v1"
	v2 "go_python_serve/app/api/v2"
	"go_python_serve/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// register socket api /api/ws/xxx
func RegisterWsGroup(router *gin.RouterGroup) *gin.RouterGroup {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.BuildOKResponse(gin.H{
			"message": "message",
			"nick":    "nick",
		}))
	})

	return router
}
