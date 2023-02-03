package api

import (
	"gin_serve/app/config"
	"gin_serve/app/dto"
	"gin_serve/app/model"
	"gin_serve/app/repo"
	"gin_serve/app/service"
	"gin_serve/app/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Register
// @Summary	Account
// @Schemes
// @Description	register
// @Tags	    example
// @Accept	    json
// @Produce		json
// @Param       user  body        dto.UserRegisterDTO  true  "UserRegisterDTO JSON"
// @Success		200	  {object}	  utils.Response
// @Failure     400   {object}    utils.Response
// @Router		/api/register [post]
func Register(ctx *gin.Context) {

	var user dto.UserRegisterDTO

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BuildErrorResponse(1, "register failed!", err.Error()))
		return
	}

	authService := service.NewAuthService(repo.NewUserRepository(config.DB))

	if duplicate := authService.IsDuplicateEmail(user.Email); duplicate {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.BuildErrorResponse(1, "register failed!", "email is exist!"))
		return
	}

	u := authService.CreateUser(user)

	ctx.JSON(http.StatusCreated, utils.BuildResponse("success", u))
}

// Login
// @Summary	Account
// @Schemes
// @Description	login
// @Tags	    example
// @Accept	    json
// @Produce		json
// @Param       user  body   	  dto.UserLoginDTO  true   "UserLoginDTO json"
// @Success		200	  {object}	  utils.Response
// @Failure     400   {object}    utils.Response
// @Router		/api/login [post]
func Login(ctx *gin.Context) {
	var user dto.UserLoginDTO

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.BuildErrorResponse(1, "fail", err.Error()))
		return
	}

	authService := service.NewAuthService(repo.NewUserRepository(config.DB))

	u, err := authService.VerifyCredential(user.Email, user.Password)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.BuildErrorResponse(1, "fail", err.Error()))
		return
	}

	token, err := utils.GenerateToken(strconv.Itoa(int(u.ID)))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.BuildErrorResponse(1, "token generate fail", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.BuildResponse("success", token))
}

// Logout
// @Summary	Account
// @Schemes
// @Description	logout
// @Tags	    example
// @Success		200	  {object}	utils.Response
// @Failure     400   {object}  utils.Response
// @Router		/api/logout [post]
func Logout(c *gin.Context) {

	c.JSON(http.StatusOK, utils.BuildResponse("success", utils.EmptyObj{}))
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
