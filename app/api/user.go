package api

import (
	"fmt"
	"gin_serve/app/dto"
	"gin_serve/app/model"
	"gin_serve/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context) {

	id := c.GetUint("id")

	c.JSON(http.StatusOK, utils.BuildResponse("success", model.User{
		ID:     id,
		Name:   "XiaoShao",
		Email:  "xiaoshaoqq@gmail.com",
		Gender: "M",
		Avatar: "/",
	}))
}

func DeleteUser(c *gin.Context) {
	// id := c.Param("id")
	id := c.GetUint("id")

	c.JSON(http.StatusOK, utils.BuildResponse("success", model.User{
		ID:     id,
		Name:   "XiaoShao",
		Email:  "xiaoshaoqq@gmail.com",
		Gender: "M",
		Avatar: "/",
	}))
}

func UpdateUser(c *gin.Context) {
	// id := c.Param("id")
	id := c.GetUint("id")

	c.JSON(http.StatusOK, utils.BuildResponse("success", model.User{
		ID:     id,
		Name:   "XiaoShao",
		Email:  "xiaoshaoqq@gmail.com",
		Gender: "M",
		Avatar: "/",
	}))
}

func CreateUser(c *gin.Context) {

	var user dto.UserCreateDTO

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(1, "register failed!", err.Error()))
		return
	}

	fmt.Println(user)

	c.JSON(http.StatusOK, utils.BuildResponse("success", ""))
}
