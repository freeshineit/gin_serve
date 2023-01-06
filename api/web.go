package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//	index page.
//
// router [/]
func IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "this is title",
	})
}

// login page
// router [/login.html]
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

// register page
// router [/register.html]
func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

// list page
// router [/list.html]
func ListPage(c *gin.Context) {
	c.HTML(http.StatusOK, "list.html", gin.H{})
}
