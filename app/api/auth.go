package api

import (
	"gin_serve/app/dto"
	"gin_serve/app/model"
	"gin_serve/app/repo"
	"gin_serve/app/service"
	"gin_serve/config"
	"gin_serve/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Register
// @Summary	Account
// @Schemes
// @Description	User register
// @Tags	    account
// @Accept	    json
// @Produce		json
// @Param       user  body        dto.UserRegisterDTO  true  "UserRegisterDTO JSON"
// @Success		200	  {object}	  helper.Response  "success"
// @Failure     400   {object}    helper.Response  "register failed!"
// @Router		/api/register [post]
func Register(ctx *gin.Context) {

	var user dto.UserRegisterDTO

	if err := ctx.ShouldBind(&user); err != nil {

		// errs, _ := err.(validator.ValidationErrors)

		// if !ok {
		// 	ctx.AbortWithStatusJSON(http.StatusBadGateway, "")
		// }

		// log.Println(errs.Translate(helper.Trans))

		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(1, "register failed!", err.Error()))
		return
	}

	authService := service.NewAuthService(repo.NewUserRepo(config.DB))

	if duplicate := authService.IsDuplicateEmail(user.Email); duplicate {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(1, "register failed!", "email is exist!"))
		return
	}

	u := authService.CreateUser(user)

	if err := helper.SendActiveEmail(&u); err != nil {
		zap.S().Fatal(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(1, "send email fail!", "send email fail!"))
		return
	}

	ctx.JSON(http.StatusCreated, helper.BuildResponse("success", u))
}

// Login
// @Summary	Account
// @Schemes
// @Description	User login
// @Tags	    account
// @Accept	    application/json
// @Produce		json
// @Param       user  body   	  dto.UserLoginDTO  true   "UserLoginDTO json"  default(dto.UserLoginDTO{"xiaoshaoqq@gmail.com", "123456"})
// @Success		200	  {object}	  helper.Response   "success"
// @Failure     400   {object}    helper.Response   "failed"
// @Router		/api/login [post]
func Login(ctx *gin.Context) {
	var userDTO dto.UserLoginDTO

	if err := ctx.ShouldBind(&userDTO); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(1, "fail", err.Error()))
		return
	}

	serviceCaptcha := service.NewCaptchaService()

	ok := serviceCaptcha.VerifyCaptcha(userDTO.CaptchaID, userDTO.Code)

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(1, "fail", "验证码错误"))
		return
	}

	authService := service.NewAuthService(repo.NewUserRepo(config.DB))

	u, err := authService.VerifyCredential(userDTO.Email, userDTO.Password)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(1, "fail", err.Error()))
		return
	}

	if *u.IsActive != 1 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(1, "fail", "用户还未激活，请激活后再试！"))
		return
	}

	token, err := helper.GenerateToken(u.ID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(1, "token generate fail", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helper.BuildResponse("success", token))
}

// Logout
// @Summary	Account
// @Schemes
// @Description	User logout
// @Tags	    account
// @Success		200	  {object}	helper.Response  "success"
// @Failure     400   {object}  helper.Response  "failed"
// @Router		/api/logout [post]
func Logout(c *gin.Context) {

	c.JSON(http.StatusOK, helper.BuildResponse("success", helper.EmptyObj{}))
}

// Refresh login token
// @Summary	Account
// @Schemes
// @Description	Refresh login token
// @Tags	    account
// @Accept	    json
// @Produce		json
// @Success		200	  {object}	  helper.Response  "success"
// @Failure     400   {object}    helper.Response  "failed"
// @Router		/api/refresh [post]
func Refresh(c *gin.Context) {
	var user model.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, helper.BuildErrorResponse(1, "use should bind error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse("success", user))
}

// Verify Email
// @Summary	Account
// @Schemes
// @Description	Verify Email
// @Tags	    account
// @Accept	    json
// @Produce		json
// @Success		200	  {object}	  helper.Response  "success"
// @Failure     400   {object}    helper.Response  "failed"
// @Router		/api/verify/{id} [post]
func VerifyEmail(c *gin.Context) {
	token := c.Param("token")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(1, "no token", "no token"))
		return
	}

	authService := service.NewAuthService(repo.NewUserRepo(config.DB))

	ok := authService.VerifyEmail(token)
	if ok {
		c.JSON(http.StatusOK, helper.BuildResponse("success", "激活成功"))
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(1, "fail", "激活失败"))
}
