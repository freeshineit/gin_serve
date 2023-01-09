package models

// ResponseResult ok response result
type ResponseOkResult[T any] struct {
	// http status code
	Code int32 `json:"code"`
	// http message
	Msg string `json:"msg"`
	// http data
	Data T `json:"data"`
}

// ResponseResult error response result
type ResponseErrorResult[E any] struct {
	// http status code
	Code int32 `json:"code"`
	// http message
	Msg string `json:"msg"`
	// http error
	Error E `json:"error"`
}

// BuildOkResponse serialize ok response data
func BuildOKResponse[T interface{}](data T) ResponseOkResult[T] {
	return ResponseOkResult[T]{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
}

// BuildErrorResponse serialize error response data
func BuildErrorResponse[E any](msg string, Err E) ResponseErrorResult[E] {
	return ResponseErrorResult[E]{
		Code:  1,
		Msg:   msg,
		Error: Err,
	}
}
