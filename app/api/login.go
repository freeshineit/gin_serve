package api

import (
	"fmt"
	"gin_serve/app/dto"
	"gin_serve/app/model"
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
// @Success		200	  {object}	  utils.Response
// @Failure     400   {object}    utils.Response
// @Router		/api/register [post]
func Register(c *gin.Context) {

	var user dto.RegisterDTO

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(1, "register failed!", err.Error()))
		return
	}

	fmt.Println(user)

	c.JSON(http.StatusCreated, utils.BuildResponse("success", ""))
}

// Login
// @Summary	Account
// @Schemes
// @Description	login
// @Tags	    example
// @Accept	    json
// @Produce		json
// @Param       id     path   int  true   "todo id"
// @Success		200	  {object}	  utils.Response
// @Failure     400   {object}    utils.Response
// @Router		/api/login [post]
func Login(c *gin.Context) {
	var user dto.LoginDTO

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(1, "use should bind error", err.Error()))
		return
	}

	token, err := utils.GenerateToken(user.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(1, "token generate fail", err.Error()))
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
// @Param       id    path   int  true   "todo id"
// @Success		200	  {object}	utils.Response
// @Failure     400   {object}  utils.Response
// @Router		/api/logout [post]
func Logout(c *gin.Context) {
	var user model.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, utils.BuildErrorResponse(1, "use should bind error", err.Error()))
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
// @Success		200	  {object}	  utils.Response
// @Failure     400   {object}    utils.Response
// @Router		/api/refresh [post]
func Refresh(c *gin.Context) {
	var user model.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, utils.BuildErrorResponse(1, "use should bind error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.BuildResponse("success", user))
}
