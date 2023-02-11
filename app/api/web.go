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
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":     "Gin Serve",
		"proxyPort": config.Conf.Proxy.Port,
	})
}

// login page
// router [/login]
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

// register page
// router [/register]
func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

// list page
// router [/list]
func ListPage(c *gin.Context) {
	c.HTML(http.StatusOK, "list.html", gin.H{})
}

// socket page
// router [/socket]
func SocketPage(c *gin.Context) {
	c.HTML(http.StatusOK, "socket.html", gin.H{
		"port": config.Conf.Socket.Port,
	})
}

func AuthPage(c *gin.Context) {
	// c.Redirect(http.StatusFound, "/login")
}
