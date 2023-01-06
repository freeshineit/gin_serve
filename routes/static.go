package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetStaticFS(r *gin.Engine) {
	// set html template
	r.LoadHTMLGlob("./templates/*.html")

	// set server static
	r.StaticFile("favicon.ico", "./public/favicon.ico")
	r.StaticFS("/static", http.Dir("public/static"))
	// r.StaticFS("/upload", http.Dir("upload"))
}
