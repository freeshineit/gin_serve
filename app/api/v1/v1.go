package v1

import (
	"gin_serve/app/utils"
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
//	@Success		200	 {object}	utils.BuildResponse
//	@Router			/v1/list [get]
func List(c *gin.Context) {
	c.JSON(http.StatusOK, utils.BuildResponse("success", gin.H{
		"message": "v1 api",
		"nick":    "v1 api",
	}))
}
