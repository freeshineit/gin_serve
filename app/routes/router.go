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
		apiGroup.POST("/file_upload", api.FileUpload)

		// // json提交
		// apiGroup.POST("/json_post", api.JSONPost)

		// //url encode 提交
		// apiGroup.POST("/urlencoded_post", api.UrlencodedPost)

		// // 即支持json又支持form
		// apiGroup.POST("/json_and_form_post", api.JSONAndFormPost)

		// // xml 提交
		// apiGroup.POST("/xml_post", api.XMLPost)
		apiGroup.GET("/query", func(c *gin.Context) {
			// message := c.Query("message")
			// nick := c.DefaultQuery("nick", "anonymous")

			c.JSON(http.StatusOK, models.BuildOKResponse(gin.H{
				"message": "message",
				"nick":    "nick",
			}))
		})
	}

	// /api/v1
	apiV1Group := router.Group("/api/v1", middleware.JwtAuth())
	{
		// // 表单提交
		// apiGroup.POST("/form_post", api.FormPost)

		// // json提交
		// apiGroup.POST("/json_post", api.JSONPost)

		// //url encode 提交
		// apiGroup.POST("/urlencoded_post", api.UrlencodedPost)

		// // 即支持json又支持form
		// apiGroup.POST("/json_and_form_post", api.JSONAndFormPost)

		// // xml 提交
		// apiGroup.POST("/xml_post", api.XMLPost)

		// // 文件上传
		// apiGroup.POST("/file_upload", api.FileUpload)

		// // 文件分片上传
		// apiGroup.POST("/file_chunk_upload", api.FileChunkUpload)

		apiV1Group.GET("/query", func(c *gin.Context) {
			// message := c.Query("message")
			// nick := c.DefaultQuery("nick", "anonymous")

			c.JSON(http.StatusOK, models.BuildOKResponse(gin.H{
				"message": "message",
				"nick":    "nick",
			}))
		})
	}

	// /api/v2
	apiV2Group := router.Group("/api/v2", middleware.JwtAuth())
	{

		// // 表单提交
		// apiGroup.POST("/form_post", api.FormPost)

		// // json提交
		// apiGroup.POST("/json_post", api.JSONPost)

		// //url encode 提交
		// apiGroup.POST("/urlencoded_post", api.UrlencodedPost)

		// // 即支持json又支持form
		// apiGroup.POST("/json_and_form_post", api.JSONAndFormPost)

		// // xml 提交
		// apiGroup.POST("/xml_post", api.XMLPost)

		// // 文件上传
		// apiGroup.POST("/file_upload", api.FileUpload)

		// // 文件分片上传
		// apiGroup.POST("/file_chunk_upload", api.FileChunkUpload)

		apiV2Group.GET("/query", func(c *gin.Context) {
			// message := c.Query("message")
			// nick := c.DefaultQuery("nick", "anonymous")

			c.JSON(http.StatusOK, models.BuildOKResponse(gin.H{
				"message": "message",
				"nick":    "nick",
			}))
		})
	}

	return router
}
