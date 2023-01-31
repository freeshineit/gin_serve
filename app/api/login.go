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
// @Success		200	{string}	models.BuildOKResponse(gin.H{"message": "v1 api","nick": "v1 api",})
// @Router		/api/register [post]
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

// Login
// @Summary	Account
// @Schemes
// @Description	login
// @Tags	    example
// @Accept	    json
// @Produce		json
// @Param       id     path   int  true   "todo id"
// @Success		200	{string}	models.BuildOKResponse(gin.H{"message": "v1 api","nick": "v1 api",})
// @Router		/api/login [post]
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

// Logout
// @Summary	Account
// @Schemes
// @Description	logout
// @Tags	    example
// @Accept	    json
// @Produce		json
// @Param       id     path   int  true   "todo id"
// @Success		200	{string}	models.BuildOKResponse(gin.H{"message": "v1 api","nick": "v1 api",})
// @Router		/api/logout [post]
func Logout(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, models.BuildErrorResponse[any]("use should bind error", err))

	}

	c.JSON(http.StatusOK, models.BuildOKResponse(user))
}

// Refresh login token
// @Summary	Account
// @Schemes
// @Description	Refresh login token
// @Tags	    example
// @Accept	    json
// @Produce		json
// @Success		200	{string}	models.BuildOKResponse(gin.H{"message": "v1 api","nick": "v1 api",})
// @Router		/api/refresh [post]
func Refresh(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, models.BuildErrorResponse[any]("use should bind error", err))
		return
	}

	c.JSON(http.StatusOK, models.BuildOKResponse(user))
}
