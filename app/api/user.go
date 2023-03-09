package api

import (
	"gin_serve/app/constant"
	"gin_serve/app/repo"
	"gin_serve/app/service"
	"gin_serve/config"
	"gin_serve/helper"
	"gin_serve/message"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// User
// @Summary	User
// @Schemes
// @Description	Get user by id
// @Tags	    account
// @Accept	    application/json
// @Produce		json
// @Param       user  path   	  int  true
// @Success		200	  {object}	  helper.Response   "success"
// @Failure     400   {object}    helper.Response   "failed"
// @Router		/api/user/{id} [get]
func GetUserByID(ctx *gin.Context) {

	id := ctx.Param("id")
	nID, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(http.StatusBadRequest, message.InvalidParameter, err.Error()))
		return
	}

	tokenClaims, exit := ctx.Get(constant.CtxTokenClaimsKey)

	if exit {
		userID := tokenClaims.(*helper.JWTAuthCustomClaim).UserID

		// 判断当前范访问的用户是否是同一个用户， 否则不允许访问
		if userID == nID {
			service := service.NewAuthService(repo.NewUserRepo(config.DB))

			t := service.FindByID(nID)

			if t.Email == "" {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(http.StatusBadRequest, "user no exists", "user no exists"))
			} else {
				ctx.JSON(http.StatusCreated, helper.BuildResponse("success", t))
			}

			return
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(http.StatusBadRequest, "no promise", "no promise"))
	}

}

// User
// @Summary	User
// @Schemes
// @Description	Get user by token
// @Tags	    account
// @Accept	    application/json
// @Produce		json
// @Success		200	  {object}	  helper.Response   "success"
// @Failure     400   {object}    helper.Response   "failed"
// @Router		/api/me [get]
func GetMe(ctx *gin.Context) {

	tokenClaims, exit := ctx.Get(constant.CtxTokenClaimsKey)

	if exit {
		service := service.NewAuthService(repo.NewUserRepo(config.DB))
		userID := tokenClaims.(*helper.JWTAuthCustomClaim).UserID
		t := service.FindByID(userID)

		if t.Email == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(http.StatusBadRequest, "user no exists", "user no exists"))
		} else {
			ctx.JSON(http.StatusOK, helper.BuildResponse("success", t))
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(http.StatusBadRequest, "user no exists", "user no exists"))
		return
	}
}

func DeleteUser(c *gin.Context) {

	// id := c.Param("id")

	// nID, err := strconv.ParseUint(id, 10, 0)

	// if err != nil {

	// }

	// c.JSON(http.StatusOK, helper.BuildResponse("success", model.User{
	// 	ID:     nID,
	// 	Name:   "XiaoShao",
	// 	Email:  "xiaoshaoqq@gmail.com",
	// 	Gender: "M",
	// 	Avatar: "/",
	// }))
}

func UpdateUser(c *gin.Context) {
	// id := c.Param("id")

	// nID, err := strconv.ParseUint(id, 10, 0)

	// if err != nil {

	// }

	// c.JSON(http.StatusOK, helper.BuildResponse("success", model.User{
	// 	ID:     nID,
	// 	Name:   "XiaoShao",
	// 	Email:  "xiaoshaoqq@gmail.com",
	// 	Gender: "M",
	// 	Avatar: "/",
	// }))
}

func CreateUser(c *gin.Context) {

	// var user dto.UserRegisterDTO

	// if err := c.ShouldBind(&user); err != nil {
	// 	c.JSON(http.StatusBadRequest, helper.BuildErrorResponse(http.StatusBadRequest, message.InvalidParameter, helper.ParseBindingError(err)))
	// 	return
	// }

	// fmt.Println(user)

	// c.JSON(http.StatusOK, helper.BuildResponse("success", ""))
}
