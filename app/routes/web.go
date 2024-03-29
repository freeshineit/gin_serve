package routes

import (
	"gin_serve/app/api"

	"github.com/gin-gonic/gin"
)

// page routes
func SetRoutesPage(r *gin.Engine) {
	// 首页 router /
	r.GET("/", api.IndexPage)
	r.GET("/login", api.LoginPage)
	r.GET("/register", api.RegisterPage)
	r.GET("/list", api.ListPage)
	r.GET("/socket", api.SocketPage)
	r.GET("/verify_email/:id", api.VerifyEmailPage)
}
