package v2

import (
	"gin_serve/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	c.JSON(http.StatusOK, helper.BuildResponse("success", gin.H{
		"message": "v2 api",
		"nick":    "v2 api",
	}))
}
