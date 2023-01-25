package v1

import (
	"gin_serve/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api
// List
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} models.BuildOKResponse(gin.H{"message": "v1 api","nick":    "v1 api",})
// @Router /v1/list [get]
func List(c *gin.Context) {
	c.JSON(http.StatusOK, models.BuildOKResponse(gin.H{
		"message": "v1 api",
		"nick":    "v1 api",
	}))
}
