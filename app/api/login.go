package api

import (
	"fmt"
	"go_python_serve/app/models"
	"go_python_serve/app/utils"
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

	token, err := utils.GenerateToken(&user)

	if err != nil {
	}

	// c.JSON(http.StatusOK, models.BuildOKResponse(user))
	c.JSON(http.StatusOK, models.BuildOKResponse(token))
}

// 登录
func Login(c *gin.Context) {
	// var user models.User

	// if err := c.ShouldBind(&user); err != nil {
	// 	c.JSON(http.StatusOK, models.BuildErrorResponse[any]("use should bind error", err))
	// }

	jwtUser, valid := utils.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjoiIiwibmFtZSI6InF3ZXEiLCJlbWFpbCI6InF3ZUBnLmNvbSIsImdlbmRlciI6IjIzIiwiYXZhdGFyIjoiaW1hZ2VzL3VwbG9hZC9mdW5jdGlvbl9jbGFzcy5wbmcifSwiaXNzIjoiU2hpbmVTaGFvIiwic3ViIjoieGlhb3NoYW9xcUBnbWFpbC5jb20iLCJhdWQiOlsiX0F1ZGllbmNlXyJdLCJleHAiOjE2NzM2Mjc5MTMsIm5iZiI6MTY3MzYyMDcxMywiaWF0IjoxNjczNjIwNzEzLCJqdGkiOiIxIn0.7fcMsVzxUpvx_GvFXkcboVzX3CrweqiPdp8j70_f6t8")

	fmt.Println(jwtUser)
	fmt.Println(valid)

	if !valid {
		c.JSON(http.StatusOK, models.BuildErrorResponse[any]("fail", "无效"))
		return
	}

	c.JSON(http.StatusOK, models.BuildOKResponse(jwtUser))
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
