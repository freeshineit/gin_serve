package middleware

import (
	"github.com/gin-gonic/gin"
)

func SetMiddleware(router *gin.Engine) {
	// cors
	Cors(router)
	// jwt

}
