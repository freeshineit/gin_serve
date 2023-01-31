package api

import (
	"gin_serve/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context) {

	id := c.GetUint("id")

	c.JSON(http.StatusOK, models.BuildOKResponse(models.User{
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

	c.JSON(http.StatusOK, models.BuildOKResponse(models.User{
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

	c.JSON(http.StatusOK, models.BuildOKResponse(models.User{
		ID:     id,
		Name:   "XiaoShao",
		Email:  "xiaoshaoqq@gmail.com",
		Gender: "M",
		Avatar: "/",
	}))
}

func CreateUser(c *gin.Context) {

	c.JSON(http.StatusOK, models.BuildOKResponse(models.User{
		ID:     9999,
		Name:   "XiaoShao",
		Email:  "xiaoshaoqq@gmail.com",
		Gender: "M",
		Avatar: "/",
	}))
}
