package v2

import (
	"go_python_serve/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	c.JSON(http.StatusOK, models.BuildOKResponse(gin.H{
		"message": "message",
		"nick":    "nick",
	}))
}
