package v1

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"gin_serve/app/models"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// func TestTodoHandle(t *testing.T) {
// 	r := gin.Default()

// 	r.POST("/api/v1/todo", CreateTodo)

// 	r.GET("/api/v1/todos", GetTodos)
// 	r.GET("/api/v1/todo/:id", GetTodo)
// 	r.PUT("/api/v1/todo/:id/content", PutTodoContent)
// 	r.PUT("/api/v1/todo/:id/status", PutTodoStatus)
// 	r.DELETE("/api/v1/todo/:id", DeleteTodo)

// 	// create
// 	todo := models.Todo{
// 		Content: "this is test",
// 	}

// 	jsonValue, _ := json.Marshal(todo)
// 	fmt.Println(bytes.NewBuffer(jsonValue))

// 	req, _ := http.NewRequest("POST", "/api/v1/todo", bytes.NewBuffer(jsonValue))

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusCreated, w.Code)
// }
