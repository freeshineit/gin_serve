package api

import (
	"gin_serve/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	index page.
//
// router [/]
func IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.html", gin.H{
		"title":     "Gin Serve",
		"proxyPort": config.Conf.Proxy.Port,
	})
}

// login page
// router [/login]
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/login.html", gin.H{
		"title": "Login",
	})
}

// register page
// router [/register]
func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/register.html", gin.H{
		"title": "Register",
	})
}

// list page
// router [/list]
func ListPage(c *gin.Context) {
	c.HTML(http.StatusOK, "list/index.html", gin.H{
		"proxyPort": config.Conf.Proxy.Port,
	})
}

// socket page
// router [/socket]
func SocketPage(c *gin.Context) {
	c.HTML(http.StatusOK, "socket/index.html", gin.H{
		"title":     "Socket",
		"port":      config.Conf.Socket.Port,
		"proxyPort": config.Conf.Proxy.Port,
	})
}

func VerifyEmailPage(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/verify_email_active.html", gin.H{
		"title": "Email Active",
		// "port":      config.Conf.Socket.Port,
		// "proxyPort": config.Conf.Proxy.Port,
	})
}
