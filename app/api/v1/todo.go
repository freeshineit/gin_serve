package v1

import (
	"gin_serve/app/models"
	"gin_serve/app/utils"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Todos = make([]models.Todo, 0)

// Get todo by id
// @Summary	Todo
// @Schemes
// @Description	do ping
// @Tags	    example
// @Accept	    json
// @Produce		json
// @Param       id   path   int  true   "todo id"
// @Success		200	 {object}	utils.Response
// @Router		/api/v1/todo/{id} [get]
func GetTodo(c *gin.Context) {
	id := c.Param("id")

	aid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusOK, utils.BuildErrorResponse(1, "fail", "id convert failed"))
		return
	}

	for _, t := range Todos {
		if t.ID == uint(aid) {
			// Todos.
			c.JSON(http.StatusOK, utils.BuildResponse("success", t))
			return
		}
	}

	c.JSON(http.StatusOK, utils.BuildErrorResponse(1, "fail", "id not exist"))
}

// Get todo list
// @Summary	Todo
// @Schemes
// @Description	Get todo list
// @Tags	    example
// @Accept		json
// @Produce		json
// @Success		200	 {object}	utils.Response
// @Router		/api/v1/todos [get]
func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, utils.BuildResponse("success", Todos))
}

// Update todo content by id
// @Summary	Todo
// @Schemes
// @Description	Update todo content by id
// @Tags		example
// @Accept		json
// @Produce		json
// @Success		200	 {object}	utils.Response
// @Router		/api/v1/todo/{id}/content [put]
func PutTodoContent(c *gin.Context) {
	id := c.Param("id")
	aid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusOK, utils.BuildErrorResponse(1, "fail", "id convert failed"))
		return
	}

	type Content struct {
		Content string `json:"content" form:"content" binding:"required"`
	}

	var con Content

	// 绑定
	if err := c.ShouldBind(&con); err != nil {
		c.JSON(http.StatusOK, utils.BuildErrorResponse(1, "fail", err.Error()))
		return
	}

	for i, t := range Todos {
		if t.ID == uint(aid) {
			// Todos.
			Todos[i].Content = con.Content
			c.JSON(http.StatusOK, utils.BuildResponse[any]("success", nil))
			return
		}
	}

	c.JSON(http.StatusOK, utils.BuildErrorResponse(1, "fail", "id not exist"))
}

// Update todo status by id
// @Summary	Todo
// @Schemes
// @Description	Update todo status by id
// @Tags		example
// @Accept		json
// @Produce		json
// @Success		200	 {object}	utils.Response
// @Router		/api/v1/todo/{id}/status [put]
func PutTodoStatus(c *gin.Context) {

	id := c.Param("id")
	aid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusOK, utils.BuildErrorResponse(1, "fail", "id convert failed"))
		return
	}

	for i, t := range Todos {
		if t.ID == uint(aid) {
			if Todos[i].Status == 0 {
				Todos[i].Status = 1
			} else if Todos[i].Status == 1 {
				Todos[i].Status = 0
			} else {
				c.JSON(http.StatusOK, utils.BuildErrorResponse(1, "update fail", "id not exist"))
				return
			}
			c.JSON(http.StatusOK, utils.BuildResponse("success", "update success"))
			return
		}
	}

	c.JSON(http.StatusOK, utils.BuildErrorResponse(1, "fail", "update fail"))
}

// Delete todo by id
// @Summary	Todo
// @Schemes
// @Description Delete todo by id
// @Tags		example
// @Accept		json
// @Produce		json
// @Success		200	 {object}  utils.Response
// @Router		/v1/todo/{id} [delete]
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	aid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusOK, utils.BuildErrorResponse(1, "fail", "id convert failed"))
		return
	}

	for i, todo := range Todos {

		if todo.ID == uint(aid) {
			// Todos.
			Todos = append(Todos[:i], Todos[i+1:]...)
			c.JSON(http.StatusOK, utils.BuildResponse("success", "delete success"))
			return
		}
	}

	c.JSON(http.StatusOK, utils.BuildErrorResponse(1, "fail", "id not exist"))
}

// Create todo
// @Summary	Todo
// @Schemes
// @Description	Create todo
// @Tags		example
// @Accept		json
// @Produce		json
// @Success		200	 {object}	utils.Response
// @Router		/api/v1/todo [post]
func CreateTodo(c *gin.Context) {

	var todo models.Todo

	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(1, "create fail", err.Error()))
		return
	}

	r := rand.New(rand.NewSource(99))
	todo.ID = uint(r.Uint32()) //

	Todos = append([]models.Todo{todo}, Todos...)

	c.JSON(http.StatusCreated, utils.BuildResponse("success", todo))
}
