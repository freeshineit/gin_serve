package helper

import (
	"gin_serve/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// Generate Token
// @Example GenerateToken(1)
func GenerateToken(UserID string) (string, error) {

	jwtConfig := config.Conf.JWT

	claims := jwtCustomClaim{
		UserID: UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 24h
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    jwtConfig.Issuer,
			Subject:   "xiaoshaoqq@gmail.com",
			ID:        "110",
			Audience:  []string{"_Audience_"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtConfig.Secret))
}

// Validate Token
// @Example ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzNDU2Nzg5MCJ9.HhclBU1hdg0RynbUgnLXtm9rhm0m4yuWJF0jjVaZ_u0")
func ValidateToken(token string) (*jwt.Token, error) {
	jwtConfig := config.Conf.JWT

	return jwt.ParseWithClaims(token, &jwtCustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConfig.Secret), nil
	})
}
