package api

import (
	"go_python_serve/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册
func Register(c *gin.Context) {
	var user models.User

	// 绑定不成功
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, models.BuildErrorResponse("fail", err))
		return
	}

	c.JSON(http.StatusOK, models.BuildOKResponse(user))
}

// 登录
func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, models.BuildErrorResponse[any]("use should bind error", err))

	}

	c.JSON(http.StatusOK, models.BuildOKResponse(user))
}

// 登出
func Logout(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, models.BuildErrorResponse[any]("use should bind error", err))

	}

	c.JSON(http.StatusOK, models.BuildOKResponse(user))
}

// token 刷新
func Refresh(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, models.BuildErrorResponse[any]("use should bind error", err))

	}

	c.JSON(http.StatusOK, models.BuildOKResponse(user))
}
