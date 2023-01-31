package utils

import "strings"

type Response[D any] struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Error   []string `json:"error"`
	Data    D        `json:"data"`
}

type EmptyObj struct{}

// build response success
// examples
// BuildResponse[any](1, "success", map[string]string{"name": "xxx"})
func BuildResponse[T interface{}](message string, data T) Response[T] {

	res := Response[T]{
		Code:    0,
		Message: message,
		Data:    data,
	}

	return res
}

// build response error
// err split `\n`
// examples
// BuildErrorResponse("用户名不对", "用户名不对\n密码不对")
func BuildErrorResponse(message string, err string) Response[any] {
	errs := strings.Split(err, "\n")

	res := Response[any]{
		Code:    1,
		Message: message,
		Error:   errs,
	}

	return res
}
