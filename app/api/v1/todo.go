package v1

import (
	"gin_serve/app/models"
	"gin_serve/app/utils"
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
// @Param       id     path   int  true   "todo id"
// @Success		200	{string}	utils.BuildResponse("success", todo)
// @Router		/api/v1/todo/{id} [get]
func GetTodo(c *gin.Context) {
	id := c.Param("id")

	id64, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		c.JSON(http.StatusOK, utils.BuildErrorResponse("fail", "id convert failed"))
		return
	}

	for _, t := range Todos {
		if t.ID == uint(id64) {
			// Todos.
			c.JSON(http.StatusOK, utils.BuildResponse("success", t))
			return
		}
	}

	c.JSON(http.StatusOK, utils.BuildErrorResponse("fail", "id not exist"))
}

// Get todo list
// @Summary	Todo
// @Schemes
// @Description	Get todo list
// @Tags	    example
// @Accept		json
// @Produce		json
// @Success		200	{string}	utils.BuildResponse("success", gin.H{"message": "v1 api","nick":    "v1 api",})
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
// @Success		200	{string}	utils.BuildResponse("success",gin.H{"message": "v1 api","nick":    "v1 api",})
// @Router		/api/v1/todo/{id}/content [put]
func PutTodoContent(c *gin.Context) {
	// id := c.Param("id")

	id := c.GetUint("id")

	type Content struct {
		Content string `json:"content" form:"content" binding:"required"`
	}

	var con Content

	// 绑定
	if err := c.ShouldBind(&con); err != nil {
		c.JSON(http.StatusOK, utils.BuildErrorResponse("fail", err.Error()))
		return
	}

	for i, t := range Todos {
		if t.ID == id {
			// Todos.
			Todos[i].Content = con.Content
			c.JSON(http.StatusOK, utils.BuildResponse[any]("success", nil))
			return
		}
	}

	c.JSON(http.StatusOK, utils.BuildErrorResponse("fail", "id not exist"))
}

// Update todo status by id
// @Summary	Todo
// @Schemes
// @Description	Update todo status by id
// @Tags		example
// @Accept		json
// @Produce		json
// @Success		200	{string}	utils.BuildResponse("success", gin.H{"message": "v1 api","nick":    "v1 api",})
// @Router		/api/v1/todo/{id}/status [put]
func PutTodoStatus(c *gin.Context) {
	id := c.GetUint("id")

	for i, t := range Todos {
		if t.ID == id {
			if Todos[i].Status == 0 {
				Todos[i].Status = 1
			} else if Todos[i].Status == 1 {
				Todos[i].Status = 0
			} else {
				c.JSON(http.StatusOK, utils.BuildErrorResponse("update fail", "id not exist"))
				return
			}
			c.JSON(http.StatusOK, utils.BuildResponse("success", "update success"))
			return
		}
	}

	c.JSON(http.StatusOK, utils.BuildErrorResponse("fail", "update fail"))
}

// Delete todo by id
// @Summary	Todo
// @Schemes
// @Description Delete todo by id
// @Tags		example
// @Accept		json
// @Produce		json
// @Success		200	{string}	utils.BuildResponse("success", "")
// @Router		/v1/todo/{id} [delete]
func DeleteTodo(c *gin.Context) {
	id := c.GetUint("id")

	for i, todo := range Todos {

		if todo.ID == id {
			// Todos.
			Todos = append(Todos[:i], Todos[i+1:]...)
			c.JSON(http.StatusOK, utils.BuildResponse("success", "delete success"))
			return
		}
	}

	c.JSON(http.StatusOK, utils.BuildErrorResponse("fail", "id not exist"))
}

// Create todo
// @Summary	Todo
// @Schemes
// @Description	Create todo
// @Tags		example
// @Accept		json
// @Produce		json
// @Success		200	{string}	utils.BuildResponse("success", todo)
// @Router		/api/v1/todo [post]
func CreateTodo(c *gin.Context) {

	var todo models.Todo

	// 绑定不成功
	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusCreated, utils.BuildErrorResponse("create fail", err.Error()))
		return
	}

	todo.ID = 1 //utils.GenTodoUuId()

	Todos = append([]models.Todo{todo}, Todos...)

	c.JSON(http.StatusCreated, utils.BuildResponse("success", todo))
}
