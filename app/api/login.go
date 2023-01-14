package api

import (
	"go_python_serve/app/models"
	"go_python_serve/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册
func Register(c *gin.Context) {
	var user models.UserLogin

	// 绑定不成功
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, models.BuildErrorResponse("fail", err))
		return
	}

	token, err := utils.GenerateToken(&user)

	if err != nil {
	}

	// c.JSON(http.StatusOK, models.BuildOKResponse(user))
	c.JSON(http.StatusOK, models.BuildOKResponse(token))
}

// 登录
func Login(c *gin.Context) {
	var user models.UserLogin

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, models.BuildErrorResponse[any]("use should bind error", err))

		return
	}

	token, err := utils.GenerateToken(&user)

	if err != nil {
		c.JSON(http.StatusOK, models.BuildErrorResponse[any]("token generate fail", err))

		return
	}

	c.JSON(http.StatusOK, models.BuildOKResponse(token))
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
