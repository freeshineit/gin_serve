package v1

import (
	"gin_serve/app/models"
	"gin_serve/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Todos = make([]models.Todo, 0)

// @BasePath /api
// List
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} models.BuildOKResponse(gin.H{"message": "v1 api","nick":    "v1 api",})
// @Router /v1/todo [get]
func GetTodo(c *gin.Context) {
	c.JSON(http.StatusOK, models.BuildOKResponse(Todos))
}

// @BasePath /api
// List
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} models.BuildOKResponse(gin.H{"message": "v1 api","nick":    "v1 api",})
// @Router /v1/todo/:id [get]
func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, models.BuildOKResponse(Todos))
}

// @BasePath /api
// List
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} models.BuildOKResponse(gin.H{"message": "v1 api","nick":    "v1 api",})
// @Router /v1/todo/:id [put]
func PutTodo(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo

	// 绑定不成功
	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusOK, models.BuildErrorResponse("fail", err))
		return
	}

	todo.Id = id

	for i, t := range Todos {

		if t.Id == todo.Id {
			// Todos.
			Todos[i] = todo
			c.JSON(http.StatusOK, models.BuildOKResponse("update success"))
			return
		}
	}

	c.JSON(http.StatusOK, models.BuildOKResponse("update fail"))
}

// @BasePath /api
// List
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} models.BuildOKResponse()
// @Router /v1/todo/:id [delete]
func DeleteTodo(c *gin.Context) {

	id := c.Param("id")

	// id_num, err := strconv.Atoi(id)

	// if err != nil {
	// 	c.JSON(http.StatusOK, models.BuildErrorResponse("fail", "id must be int"))
	// 	return
	// }

	for i, todo := range Todos {

		if todo.Id == id {
			// Todos.
			Todos = append(Todos[:i], Todos[i+1:]...)
			c.JSON(http.StatusOK, models.BuildOKResponse(""))
			return
		}
	}

	c.JSON(http.StatusOK, models.BuildErrorResponse("fail", "id: not exist"))
}

// @BasePath /api
// create todo
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} models.BuildOKResponse(todo)
// @Router /v1/todo [post]
func CreateTodo(c *gin.Context) {

	var todo models.Todo

	// 绑定不成功
	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusOK, models.BuildErrorResponse("create fail", err))
		return
	}

	todo.Id = utils.GenTodoUuId()

	Todos = append(Todos, todo)

	c.JSON(http.StatusOK, models.BuildOKResponse(todo))
}