package helper

import (
	"gin_serve/message"
	"strings"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   []string    `json:"error,omitempty"`
}

type EmptyObj struct{}

// build response success
// examples
// BuildResponse[any]("success", map[string]string{"name": "xxx"})
func BuildResponse[T any](msg string, data T) Response {

	res := Response{
		Code:    message.SuccessCode,
		Message: msg,
		Data:    data,
	}

	res.Code = 0

	return res
}

// build response error
// err split `\n`
// examples
// BuildErrorResponse(1, "用户名不对", "用户名不对\n密码不对")
func BuildErrorResponse(code int, msg string, err string) Response {
	errs := strings.Split(err, "\n")

	res := Response{
		Code:    code,
		Message: msg,
		Error:   errs,
	}

	return res
}
