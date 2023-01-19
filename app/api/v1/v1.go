package v1

import (
	"go_python_serve/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	c.JSON(http.StatusOK, models.BuildOKResponse(gin.H{
		"message": "v1 api",
		"nick":    "v1 api",
	}))
}
