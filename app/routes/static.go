package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetStaticFS(r *gin.Engine) {
	// set html template
	r.LoadHTMLGlob("./templates/**/*.html")
	// r.HTMLRender = helper.LoadTemplateFiles("templates", ".html")

	// set server static
	r.StaticFile("favicon.ico", "./public/favicon.ico")
	r.StaticFS("/static", http.Dir("public/static"))
	// upload file
	r.StaticFS("/images/upload", http.Dir("upload"))
}
