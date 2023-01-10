package utils

import (
	"fmt"
	"go_python_serve/app/models"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

// 生成token
func GenerateToken() (string, error) {

	secret := viper.GetString("jwt.secret")

	hmacSampleSecret := []byte(secret)
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = models.AuthUser{}

	return token.SignedString(hmacSampleSecret)
}

// token 解析
func ParseToken(tokenStr string) (*models.AuthUser, error) {
	secret := viper.GetString("jwt.secret")

	token, err := jwt.ParseWithClaims(tokenStr, &models.AuthUser{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(*models.AuthUser); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
