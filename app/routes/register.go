package routes

import (
	"go_python_serve/app/api"
	"go_python_serve/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// register with auth api /api/xxx
func RegisterWithAuthGroup(router *gin.RouterGroup) *gin.RouterGroup {
	router.POST("/file_upload", api.FileUpload)

	return router
}

// register v1 api /api/v1/xxx
func RegisterV1Group(router *gin.RouterGroup) *gin.RouterGroup {

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

// register v2 api /api/v2/xxx
func RegisterV2Group(router *gin.RouterGroup) *gin.RouterGroup {

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
