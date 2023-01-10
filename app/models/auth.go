package models

import (
	"github.com/golang-jwt/jwt"
)

type AuthUser struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Avatar string `json:"avatar"`
	Email  string `json:"email"`
	*jwt.StandardClaims
}
