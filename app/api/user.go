package api

import (
	"fmt"
	"gin_serve/app/dto"
	"gin_serve/app/model"
	"gin_serve/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context) {

	id := c.Param("id")

	nID, err := strconv.ParseUint(id, 10, 0)

	if err != nil {

	}

	c.JSON(http.StatusOK, helper.BuildResponse("success", model.User{
		ID:     nID,
		Name:   "XiaoShao",
		Email:  "xiaoshaoqq@gmail.com",
		Gender: "M",
		Avatar: "/",
	}))
}

func DeleteUser(c *gin.Context) {

	id := c.Param("id")

	nID, err := strconv.ParseUint(id, 10, 0)

	if err != nil {

	}

	c.JSON(http.StatusOK, helper.BuildResponse("success", model.User{
		ID:     nID,
		Name:   "XiaoShao",
		Email:  "xiaoshaoqq@gmail.com",
		Gender: "M",
		Avatar: "/",
	}))
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	nID, err := strconv.ParseUint(id, 10, 0)

	if err != nil {

	}

	c.JSON(http.StatusOK, helper.BuildResponse("success", model.User{
		ID:     nID,
		Name:   "XiaoShao",
		Email:  "xiaoshaoqq@gmail.com",
		Gender: "M",
		Avatar: "/",
	}))
}

func CreateUser(c *gin.Context) {

	var user dto.UserRegisterDTO

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, helper.BuildErrorResponse(1, "register failed!", err.Error()))
		return
	}

	fmt.Println(user)

	c.JSON(http.StatusOK, helper.BuildResponse("success", ""))
}
