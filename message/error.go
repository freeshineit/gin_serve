package message

import "net/http"

const (
	// Unauthorized
	UnauthorizedCode    = http.StatusUnauthorized // 402
	Unauthorized        = "no permission"         // no permission
	UnauthorizedExpired = "token is expired"      // token is expired
	// bad
	BadRequestCode = http.StatusBadRequest // 400
)
