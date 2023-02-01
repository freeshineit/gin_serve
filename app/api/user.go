package api

import (
	"fmt"
	"gin_serve/app/models"
	"gin_serve/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context) {

	id := c.GetUint("id")

	c.JSON(http.StatusOK, utils.BuildResponse("success", models.User{
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

	c.JSON(http.StatusOK, utils.BuildResponse("success", models.User{
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

	c.JSON(http.StatusOK, utils.BuildResponse("success", models.User{
		ID:     id,
		Name:   "XiaoShao",
		Email:  "xiaoshaoqq@gmail.com",
		Gender: "M",
		Avatar: "/",
	}))
}

func CreateUser(c *gin.Context) {

	var user models.UserRegister

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(1, "register failed!", err.Error()))
		return
	}

	fmt.Println(user)

	c.JSON(http.StatusOK, utils.BuildResponse("success", ""))
}
