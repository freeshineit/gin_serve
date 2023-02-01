package api

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

// func TestLoginHandle(t *testing.T) {
// 	r := gin.Default()

// 	r.POST("/api/login", Login)

// 	// create
// 	user := models.UserLogin{
// 		Email:    "xiaoshaoqq@gmail.com",
// 		Password: "123413243",
// 	}

// 	jsonValue, _ := json.Marshal(user)

// 	fmt.Println(bytes.NewBuffer(jsonValue))

// 	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonValue))

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	fmt.Println(w.Body.String())
// }
