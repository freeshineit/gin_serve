package v2

import (
	"gin_serve/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	c.JSON(http.StatusOK, utils.BuildResponse("success", gin.H{
		"message": "v2 api",
		"nick":    "v2 api",
	}))
}
