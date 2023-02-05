package v1

import (
	"gin_serve/app/dto"
	"gin_serve/app/middleware"
	"gin_serve/app/model"
	"gin_serve/app/repo"
	"gin_serve/app/service"
	"gin_serve/config"
	"gin_serve/helper"
	"gin_serve/message"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// var Todos = make([]model.Todo, 0)

// Create todo
// @Summary	Todo
// @Schemes
// @Description	Create todo
// @Tags		example
// @Accept		json
// @Produce		json
// @Success		200	 {object}	helper.Response
// @Router		/api/v1/todo [post]
// @Security    ApiKeyAuth
func CreateTodo(ctx *gin.Context) {

	var todo dto.TodoCreateDTO

	if err := ctx.ShouldBind(&todo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(1, "create fail", err.Error()))
		return
	}

	tokenClaims, exists := ctx.Get(middleware.TokenClaims)

	if exists {
		service := service.NewTodoService(repo.NewTodoRepo(config.DB))
		todo.UserID = tokenClaims.(*helper.TokenClaim).UserID
		res := service.CreateTodo(todo)
		ctx.JSON(http.StatusCreated, helper.BuildResponse("success", res))
		return
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(1, "user no exists", "user no exists"))
	}
}

// Get todo by id
// @Summary	Todo
// @Schemes
// @Description	do ping
// @Tags	    example
// @Accept	    json
// @Produce		json
// @Param       id   path   int  true   "todo id"
// @Success		200	 {object}	helper.Response
// @Router		/api/v1/todo/{id} [get]
// @Security    ApiKeyAuth
func GetTodo(c *gin.Context) {
	id := c.Param("id")
	aid, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(message.BadRequestCode, "fail", "id convert failed"))
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
// @Tags	    example
// @Accept		json
// @Produce		json
// @Param       offset query   int  true   "offset"
// @Param       page   query   int  true   "page"
// @Success		200	 {object}	helper.Response
// @Router		/api/v1/todos [get]
// @Security    ApiKeyAuth
func GetTodos(ctx *gin.Context) {
	tokenClaims, exists := ctx.Get(middleware.TokenClaims)

	if exists {
		service := service.NewTodoService(repo.NewTodoRepo(config.DB))
		userID := tokenClaims.(*helper.TokenClaim).UserID

		query := dto.PaginationRequestDTO{}
		err := ctx.ShouldBindQuery(&query)

		if err != nil {
			log.Println(err.Error())
		}

		list, total, _ := service.FindAll(userID, query.Offset, query.Page, 10)

		ctx.JSON(http.StatusCreated, helper.BuildResponse("success", dto.ListDTO[model.Todo]{
			List: list,
			Page: dto.PaginationResponseDTO{
				Offset: 1,
				Page:   1,
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
// @Tags		example
// @Accept		json
// @Produce		json
// @Param       id       path   int  true   "todo id"
// @Param       content  body   string  true   "todo content"
// @Success		200	 {object}	helper.Response
// @Router		/api/v1/todo/{id}/content [put]
// @Security    ApiKeyAuth
func PutTodoContent(ctx *gin.Context) {
	id := ctx.Param("id")
	tid, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, helper.BuildErrorResponse(1, "fail", "id convert failed"))
		return
	}

	var todoUpdateContentDTO dto.TodoUpdateContentDTO

	if err := ctx.ShouldBind(&todoUpdateContentDTO); err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, helper.BuildErrorResponse(1, "fail", "content format is incorrect"))
		return
	}

	tokenClaims, exists := ctx.Get(middleware.TokenClaims)

	if exists {
		service := service.NewTodoService(repo.NewTodoRepo(config.DB))

		userID := tokenClaims.(*helper.TokenClaim).UserID

		ok, err := service.UpdateTodoContent(tid, todoUpdateContentDTO.Content, userID)

		if ok {
			ctx.JSON(http.StatusOK, helper.BuildResponse("success", "update success"))
		} else {
			ctx.JSON(http.StatusOK, helper.BuildErrorResponse(1, "fail", err.Error()))
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
// @Tags		example
// @Accept		json
// @Produce		json
// @Param       id       path   int  true   "todo id"
// @Param       status   body   int  true   "todo status"
// @Success		200	 {object}	helper.Response
// @Router		/api/v1/todo/{id}/status [put]
// @Security    ApiKeyAuth
func PutTodoStatus(ctx *gin.Context) {

	id := ctx.Param("id")
	tid, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, helper.BuildErrorResponse(1, "fail", "id convert failed"))
		return
	}

	var todoUpdateStatusDTO dto.TodoUpdateStatusDTO

	if err := ctx.ShouldBind(&todoUpdateStatusDTO); err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, helper.BuildErrorResponse(1, "fail", "status format is incorrect"))
		return
	}

	tokenClaims, exists := ctx.Get(middleware.TokenClaims)

	if exists {
		service := service.NewTodoService(repo.NewTodoRepo(config.DB))

		userID := tokenClaims.(*helper.TokenClaim).UserID

		ok, err := service.UpdateTodoStatus(tid, *todoUpdateStatusDTO.Status, userID)

		if ok {
			ctx.JSON(http.StatusOK, helper.BuildResponse("success", "update success"))
		} else {
			ctx.JSON(http.StatusOK, helper.BuildErrorResponse(1, "fail", err.Error()))
		}
		return
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(1, "user no exists", "user no exists"))
	}
}

// Delete todo by id
// @Summary	Todo
// @Schemes
// @Description Delete todo by id
// @Tags		example
// @Accept		json
// @Produce		json
// @Param       id   path   int  true   "todo id"
// @Success		200	 {object}  helper.Response
// @Router		/v1/todo/{id} [delete]
// @Security    ApiKeyAuth
func DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	tid, err := strconv.ParseUint(id, 10, 0)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, helper.BuildErrorResponse(message.BadRequestCode, "fail", "id convert failed"))
		return
	}

	tokenClaims, exists := ctx.Get(middleware.TokenClaims)

	if exists {
		service := service.NewTodoService(repo.NewTodoRepo(config.DB))

		userID := tokenClaims.(*helper.TokenClaim).UserID

		ok, err := service.DeleteTodo(tid, userID)

		if ok {
			ctx.JSON(http.StatusOK, helper.BuildResponse("success", "delete success"))
		} else {
			ctx.JSON(http.StatusOK, helper.BuildErrorResponse(1, "fail", err.Error()))
		}
		return
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BuildErrorResponse(1, "user no exists", "user no exists"))
	}

}
