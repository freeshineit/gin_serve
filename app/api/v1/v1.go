package v1

import (
	"gin_server/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	c.JSON(http.StatusOK, models.BuildOKResponse(gin.H{
		"message": "v1 api",
		"nick":    "v1 api",
	}))
}
