package helper

import (
	"gin_serve/config"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mashingan/smapping"
)

type jwtCustomClaim struct {
	UserID uint64 `json:"user_id"`
	jwt.RegisteredClaims
}

type TokenClaim struct {
	UserID uint64 `json:"user_id"`
}

// Generate Token
// @Example GenerateToken(1)
func GenerateToken(UserID uint64) (string, error) {

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
func ValidateTokenAndClaims(tokenStr string) (*jwt.Token, *TokenClaim, error) {
	jwtConfig := config.Conf.JWT

	claim := &jwtCustomClaim{}
	token, err := jwt.ParseWithClaims(tokenStr, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConfig.Secret), nil
	})

	tokenClaim := &TokenClaim{}

	if err != nil {
		return token, tokenClaim, err
	}

	err = smapping.FillStruct(&tokenClaim, smapping.MapFields(&claim))

	return token, tokenClaim, err
}

func GetHeaderToken(ctx *gin.Context) string {
	authorization := ctx.GetHeader("Authorization")
	authorization = strings.Replace(authorization, "Basic ", "", 1)
	authorization = strings.Replace(authorization, "Bearer ", "", 1)
	return authorization
}
