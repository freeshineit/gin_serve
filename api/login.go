package api

import (
	"go_python_serve/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User

	// 绑定不成功
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, models.BuildResponse[any](http.StatusBadRequest, "fail", err))
		return
	}

	c.JSON(http.StatusOK, models.BuildResponse(http.StatusOK, "success", user))
}

func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, models.BuildResponse[any](http.StatusBadRequest, "fail", nil))

	}

	c.JSON(http.StatusOK, models.BuildResponse(http.StatusOK, "success", user))
}

func Logout(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, models.BuildResponse[any](http.StatusBadRequest, "fail", nil))

	}

	c.JSON(http.StatusOK, models.BuildResponse(http.StatusOK, "success", user))
}
