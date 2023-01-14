package utils

import (
	"fmt"
	"go_python_serve/app/models"
	"testing"

	"github.com/golang-jwt/jwt/v4"
)

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.RegisteredClaims
}

// Example creating a token using a custom claims type. The RegisteredClaims is embedded
// in the custom type to allow for easy encoding, parsing and validation of registered claims.
func TestExampleNewWithClaims(t *testing.T) {
	ss, err := GenerateToken(&models.UserLogin{
		Password: "123456",
		Email:    "xiaoshaoqq@gmail.com",
	})

	fmt.Printf("%v %v", ss, err)
	//Output: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJpc3MiOiJ0ZXN0IiwiZXhwIjoxNTE2MjM5MDIyfQ.xVuY2FZ_MRXMIEgVQ7J-TFtaucVFRXUzHm9LmV41goM <nil>
}

// Example creating a token using a custom claims type.  The StandardClaim is embedded
// in the custom type to allow for easy encoding, parsing and validation of standard claims.
func TestExampleParseWithClaims(t *testing.T) {
	// tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjoiIiwibmFtZSI6IlNoaW5lU2hhbyIsImVtYWlsIjoieGlhb3NoYW9xcUBnbWFpbC5jb20iLCJnZW5kZXIiOiJNYW4iLCJhdmF0YXIiOiIvdXBsb2FkL2F2YXRhci5wbmcifSwiaXNzIjoiU2hpbmVTaGFvIiwic3ViIjoieGlhb3NoYW9xcUBnbWFpbC5jb20iLCJhdWQiOlsiX0F1ZGllbmNlXyJdLCJleHAiOjE2NzM3MDk2MzEsIm5iZiI6MTY3MzcwMjQzMSwiaWF0IjoxNjczNzAyNDMxLCJqdGkiOiIxIn0.pXTxq-KsYgWTUtKkE8SjMqpEqEiQNSj-JFt3UEFHM7A"

	tokenString, err := GenerateToken(&models.UserLogin{
		Password: "123456",
		Email:    "xiaoshaoqq@gmail.com",
	})

	if err != nil {

	}

	user, ok := ParseToken(tokenString)
	if !ok {
		t.Errorf("ParseToken: %v %v", user, ok)
	}
}
