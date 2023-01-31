package api

import (
	"gin_serve/app/models"
	"gin_serve/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register
// @Summary	Account
// @Schemes
// @Description	register
// @Tags	    example
// @Accept	    json
// @Produce		json
// @Success		200	{object}	utils.BuildResponse("success", "token xxxxx....")
// @Failure     1   {object}    utils.BuildErrorResponse "fail" "登录失败!!"
// @Router		/api/register [post]
func Register(c *gin.Context) {
	var user models.UserLogin

	// 绑定不成功
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, utils.BuildErrorResponse("fail", err.Error()))
		return
	}

	token, err := utils.GenerateToken(&user)

	if err != nil {
		c.JSON(http.StatusOK, utils.BuildErrorResponse("fail", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponse("success", token))
}

// Login
// @Summary	Account
// @Schemes
// @Description	login
// @Tags	    example
// @Accept	    json
// @Produce		json
// @Param       id     path   int  true   "todo id"
// @Success		200	{string}	utils.BuildResponse("success", "token .....")
// @Router		/api/login [post]
func Login(c *gin.Context) {
	var user models.UserLogin

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, utils.BuildErrorResponse("use should bind error", err.Error()))
		return
	}

	token, err := utils.GenerateToken(&user)

	if err != nil {
		c.JSON(http.StatusOK, utils.BuildErrorResponse("token generate fail", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponse("success", token))
}

// Logout
// @Summary	Account
// @Schemes
// @Description	logout
// @Tags	    example
// @Accept	    json
// @Produce		json
// @Param       id     path   int  true   "todo id"
// @Success		200	{string}	utils.BuildResponse("success", nil)
// @Router		/api/logout [post]
func Logout(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, utils.BuildErrorResponse("use should bind error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponse("success", user))
}

// Refresh login token
// @Summary	Account
// @Schemes
// @Description	Refresh login token
// @Tags	    example
// @Accept	    json
// @Produce		json
// @Success		200	{string}	utils.BuildResponse("success", user)
// @Router		/api/refresh [post]
func Refresh(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, utils.BuildErrorResponse("use should bind error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponse("success", user))
}
