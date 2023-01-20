package utils

import (
	"fmt"
	"gin_server/app/config"
	"gin_server/app/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtAuthClaim struct {
	User *models.UserLogin
	jwt.RegisteredClaims
}

const TokenExpireDuration = 2 * time.Hour //过期时间

// GenerateToken generate jwt token.
func GenerateToken(u *models.UserLogin) (string, error) {

	// fmt.Print(u)

	claims := JwtAuthClaim{
		User: u,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "ShineShao",
			Subject:   "xiaoshaoqq@gmail.com",
			ID:        "1",
			Audience:  []string{"_Audience_"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.Secret))
}

// ParseToken parse jwt token
func ParseToken(tokenStr string) (user *models.UserLogin, Valid bool) {

	token, err := jwt.ParseWithClaims(tokenStr, &JwtAuthClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	})

	if err != nil {
		// fmt.Println(err)
		return nil, false
	}

	claims, ok := token.Claims.(*JwtAuthClaim)

	fmt.Println(claims, claims.User, ok)

	if ok && token.Valid {
		return claims.User, token.Valid
		// fmt.Printf("%v %v", claims.Foo, claims.RegisteredClaims.Issuer)
	} else {
		// fmt.Println(err)
		return nil, false
	}
}

// func
