package v1

import (
	"gin_serve/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@BasePath	/api
//
// List
//
//	@Summary	Test
//	@Schemes
//	@Description	do ping
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	 {object}	helper.Response
//	@Router			/v1/list [get]
func List(c *gin.Context) {
	c.JSON(http.StatusOK, helper.BuildResponse("success", gin.H{
		"message": "v1 api",
		"nick":    "v1 api",
	}))
}
