package api

import (
	"gin_serve/app/dto"
	"gin_serve/app/service"
	"gin_serve/helper"
	"net/http"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Get Captcha
// @Summary	获取验证码数据
// @Schemes
// @Description	获取验证码数据
// @Tags	    Captcha
// @Accept		json
// @Produce		json
// @Success		200	 {object}	helper.Response
// @Router		 /api/captcha [get]
func GetCaptcha(c *gin.Context) {
	length := 4
	captchaID := captcha.NewLen(length)
	var captcha dto.CaptchaDTO

	captcha.CaptchaID = captchaID
	captcha.ImageURL = "/api/captcha/" + captchaID + ".png"
	c.JSON(http.StatusOK, helper.BuildResponse("success", captcha))
}

// 验证码图片服务
// @Summary	验证码图片服务
// @Schemes
// @Description	验证码图片服务
// @Tags	    Captcha
// @Accept		json
// @Produce		json
// @Param       captchaId   path   string  true   "captcha png/wav"
// @Router		/api/captcha/{captchaId} [get]
func CaptchaServeHTTP(c *gin.Context) {
	captchaID := c.Param("captchaId")
	zap.S().Infof("GetCaptchaPng : " + captchaID)
	serviceCaptcha := service.NewCaptchaService()
	serviceCaptcha.ServerHTTP(c.Writer, c.Request)
}
