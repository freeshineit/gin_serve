package v1

import (
	"fmt"
	"gin_serve/app/constant"
	"gin_serve/app/dto"
	"gin_serve/app/repo"
	"gin_serve/app/service"
	"gin_serve/config"
	"gin_serve/helper"
	"gin_serve/message"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// var Todos = make([]model.Todo, 0)

// Create todo
// @Summary	Todo
// @Schemes
// @Description	Create todo
// @Tags		Todo
// @Accept		json
// @Produce		json
// @Param		todo body dto.TodoCreateDTO true "dto.TodoCreateDTO json"
// @Success		200	 {object}	helper.Response
// @Router		/api/v1/todo [post]
// @Security    Bearer
func CreateTodo(ctx *gin.Context) {

	var todo dto.TodoCreateDTO

	if err := ctx.ShouldBind(&todo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(http.StatusBadRequest, message.InvalidParameter, helper.ParseBindingError(err)))
		return
	}

	tokenClaims, exists := ctx.Get(constant.CtxTokenClaimsKey)

	if exists {
		service := service.NewTodoService(repo.NewTodoRepo(config.DB))
		userID := tokenClaims.(*helper.JWTAuthCustomClaim).UserID
		t, err := service.CreateTodo(todo, userID)
		if err == nil {
			ctx.JSON(http.StatusCreated, helper.BuildResponse("success", t))
			return
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(http.StatusBadRequest, "insert fail", err.Error()))
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(http.StatusBadRequest, "user no exists", "user no exists"))
	}
}

// Get todo by id
// @Summary	Todo
// @Schemes
// @Description	Get todo by id
// @Tags	    Todo
// @Accept	    json
// @Produce		json
// @Param       id   path   int  true   "todo id"
// @Success		200	 {object}	helper.Response
// @Router		/api/v1/todo/{id} [get]
// @Security    Bearer
func GetTodo(c *gin.Context) {
	id := c.Param("id")
	aid, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(message.BadRequestCode, message.InvalidParameter, "id convert failed"))
		return
	}

	service := service.NewTodoService(repo.NewTodoRepo(config.DB))

	todo := service.FindById(aid)

	c.JSON(http.StatusOK, helper.BuildResponse("success", todo))
}

// Get todo list
// @Summary	Todo
// @Schemes
// @Description	Get todo list
// @Tags	    Todo
// @Accept		json
// @Produce		json
// @Param       offset query   int  false   "offset" default(10)
// @Param       page   query   int   false  "page"   default(1)
// @Success		200	  {object}	helper.Response
// @Router		/api/v1/todos [get]
// @Security    Bearer
func GetTodos(ctx *gin.Context) {
	tokenClaims, exists := ctx.Get(constant.CtxTokenClaimsKey)

	if exists {
		service := service.NewTodoService(repo.NewTodoRepo(config.DB))
		userID := tokenClaims.(*helper.JWTAuthCustomClaim).UserID

		query := dto.PaginationRequestDTO{}
		if err := ctx.ShouldBindQuery(&query); err != nil {
			zap.S().Error(err.Error())
		}

		if query.Offset <= 0 || query.Page <= 0 {
			zap.S().Errorf("query error: %v \n", query)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(message.BadRequestCode, "fail", fmt.Sprintf("query error: %v", query)))
			return
		}

		list, total, _ := service.FindAll(userID, query.Offset, query.Page)

		ctx.JSON(http.StatusOK, helper.BuildResponse("success", dto.ListDTO[dto.TodoDTO]{
			List: list,
			Page: dto.PaginationResponseDTO{
				Offset: query.Offset,
				Page:   query.Page,
				Total:  total,
			},
		}))

		return
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(message.BadRequestCode, "fail", ""))
	}

}

// Update todo content by id
// @Summary	Todo
// @Schemes
// @Description	Update todo content by id
// @Tags		Todo
// @Accept		json
// @Produce		json
// @Param       id       path   int  true   "todo id"
// @Param       content  body   string  true   "todo content"
// @Success		200	 {object}	helper.Response
// @Router		/api/v1/todo/{id}/content [put]
// @Security    Bearer
func PutTodoContent(ctx *gin.Context) {
	id := ctx.Param("id")
	tid, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, helper.BuildErrorResponse(1, message.InvalidParameter, "id convert failed"))
		return
	}

	var todoUpdateContentDTO dto.TodoUpdateContentDTO

	if err := ctx.ShouldBind(&todoUpdateContentDTO); err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, helper.BuildErrorResponse(1, message.InvalidParameter, helper.ParseBindingError(err)))
		return
	}

	tokenClaims, exists := ctx.Get(constant.CtxTokenClaimsKey)

	if exists {
		service := service.NewTodoService(repo.NewTodoRepo(config.DB))

		userID := tokenClaims.(*helper.JWTAuthCustomClaim).UserID

		ok, err := service.UpdateTodoContent(tid, todoUpdateContentDTO.Content, userID)

		if ok {
			ctx.JSON(http.StatusOK, helper.BuildResponse("success", "update success"))
		} else {
			ctx.JSON(http.StatusOK, helper.BuildErrorResponse(http.StatusBadRequest, "fail", err.Error()))
		}
		return
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(1, "user no exists", "user no exists"))
	}
}

// Update todo status by id
// @Summary	Todo
// @Schemes
// @Description	Update todo status by id
// @Tags		Todo
// @Accept		json
// @Produce		json
// @Param       id       path   int  true   "todo id"
// @Param       status   body   int  true   "todo status"
// @Success		200	 {object}	helper.Response
// @Router		/api/v1/todo/{id}/status [put]
// @Security    Bearer
func PutTodoStatus(ctx *gin.Context) {

	id := ctx.Param("id")
	tid, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(http.StatusBadRequest, message.InvalidParameter, "id convert failed"))
		return
	}

	var todoUpdateStatusDTO dto.TodoUpdateStatusDTO

	if err := ctx.ShouldBind(&todoUpdateStatusDTO); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(1, message.InvalidParameter, helper.ParseBindingError(err)))
		return
	}

	tokenClaims, exists := ctx.Get(constant.CtxTokenClaimsKey)

	if exists {
		service := service.NewTodoService(repo.NewTodoRepo(config.DB))

		userID := tokenClaims.(*helper.JWTAuthCustomClaim).UserID

		ok, err := service.UpdateTodoStatus(tid, *todoUpdateStatusDTO.Status, userID)

		if ok {
			ctx.JSON(http.StatusOK, helper.BuildResponse("success", "update success"))
		} else {
			ctx.JSON(http.StatusBadRequest, helper.BuildErrorResponse(http.StatusBadRequest, "fail", err.Error()))
		}
		return
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(http.StatusBadRequest, "user no exists", "user no exists"))
	}
}

// Delete todo by id
// @Summary	Todo
// @Schemes
// @Description Delete todo by id
// @Tags		Todo
// @Accept		json
// @Produce		json
// @Param       id   path   int  true   "todo id"
// @Success		200	 {object}  helper.Response
// @Router		/v1/todo/{id} [delete]
// @Security    Bearer
func DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	tid, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, helper.BuildErrorResponse(message.BadRequestCode, message.InvalidParameter, "id convert failed"))
		return
	}

	tokenClaims, exists := ctx.Get(constant.CtxTokenClaimsKey)

	if exists {
		service := service.NewTodoService(repo.NewTodoRepo(config.DB))

		userID := tokenClaims.(*helper.JWTAuthCustomClaim).UserID

		ok, err := service.DeleteTodo(tid, userID)

		if ok {
			ctx.JSON(http.StatusOK, helper.BuildResponse("success", "delete success"))
		} else {
			ctx.JSON(http.StatusBadRequest, helper.BuildErrorResponse(http.StatusBadRequest, message.InvalidParameter, err.Error()))
		}
		return
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(http.StatusBadRequest, "user no exists", "user no exists"))
	}

}
