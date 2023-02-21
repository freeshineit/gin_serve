package helper

import (
	"fmt"
	"gin_serve/config"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type JWTAuthCustomClaim struct {
	UserID uint64 `json:"user_id"`
	jwt.RegisteredClaims
}

// Generate Token
// @Example GenerateToken(1)
func GenerateToken(UserID uint64) (string, error) {

	jwtConfig := config.Conf.JWT

	seconds := jwtConfig.JWTExpires

	if seconds == 0 {
		seconds = 24 * 3600 // 24h
	}

	claims := JWTAuthCustomClaim{
		UserID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(seconds))),
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

// Validate token and back claims
// @Example ValidateTokenAndClaims("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzNDU2Nzg5MCJ9.HhclBU1hdg0RynbUgnLXtm9rhm0m4yuWJF0jjVaZ_u0")
func ValidateTokenAndBackClaims(tokenStr string) (*JWTAuthCustomClaim, bool, error) {
	jwtConfig := config.Conf.JWT

	claims := &JWTAuthCustomClaim{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConfig.Secret), nil
	})

	fmt.Printf("ValidateTokenAndBackClaims: %v, %v \n", token.Valid, claims)

	return claims, token.Valid, err
}

func GetHeaderToken(ctx *gin.Context) string {
	authorization := ctx.GetHeader("Authorization")
	authorization = strings.Replace(authorization, "Basic ", "", 1)
	authorization = strings.Replace(authorization, "Bearer ", "", 1)
	return authorization
}

type JWTEmailClaim struct {
	UserID uint64 `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// Generate Token
// @Example GenerateToken(1)
func GenerateEmailToken(UserID uint64, Email string) (string, error) {

	jwtConfig := config.Conf.JWT

	claims := JWTEmailClaim{
		UserID,
		Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)), // 2h
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

// Validate email token and back claims
// @Example ValidateEmailTokenAndBackClaims("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzNDU2Nzg5MCJ9.HhclBU1hdg0RynbUgnLXtm9rhm0m4yuWJF0jjVaZ_u0")
func ValidateEmailTokenAndBackClaims(tokenStr string) (*JWTEmailClaim, bool, error) {
	jwtConfig := config.Conf.JWT

	claims := &JWTEmailClaim{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConfig.Secret), nil
	})
	return claims, token.Valid, err
}
