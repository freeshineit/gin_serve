package routes

import (
	"gin_serve/app/api"
	"gin_serve/app/middleware"
	"gin_serve/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// register with auth api /api/xxx
func RegisterGroup(router *gin.RouterGroup) *gin.RouterGroup {
	router.POST("/register", api.Register)
	router.POST("/verify_email/:token", api.VerifyEmail)

	// curl -X POST http://localhost:8080/api/login -H 'accept: application/json' -H 'Content-Type: application/json' -d '{"email": "xiaoshaoqq@gmail.com","password": "123456"}'
	router.POST("/login", api.Login)
	router.POST("/logout", api.Logout)
	router.POST("/upload", api.FileUpload) // /api/upload [post]

	router.GET("/captcha", api.GetCaptcha)                  // /api/v1/captcha [get]
	router.GET("/captcha/:captchaId", api.CaptchaServeHTTP) // /api/v1//captcha/:captchaId [get]

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
	authRouter := router.Group("/", middleware.JwtAuth())

	authRouter.POST("/upload_file", api.FileUpload) // /api/upload_file [post]
	authRouter.GET("/user/:id", api.GetUserByID)    // /api/user/:id [get]
	authRouter.POST("/user", api.CreateUser)        // /api/user [post]
	authRouter.PUT("/user/:id", api.UpdateUser)     // /api/user/:id [put]
	authRouter.DELETE("/user/:id", api.DeleteUser)  // /api/user/:id [delete]

	return authRouter
}
