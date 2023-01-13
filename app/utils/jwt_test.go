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
	ss, err := GenerateToken(&models.User{
		Name:   "string name",
		Gender: "string gender",
		Avatar: "string avatar",
		Email:  "string email",
	})

	fmt.Printf("%v %v", ss, err)
	//Output: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJpc3MiOiJ0ZXN0IiwiZXhwIjoxNTE2MjM5MDIyfQ.xVuY2FZ_MRXMIEgVQ7J-TFtaucVFRXUzHm9LmV41goM <nil>
}

// Example creating a token using a custom claims type.  The StandardClaim is embedded
// in the custom type to allow for easy encoding, parsing and validation of standard claims.
func TestExampleParseWithClaims(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjoiIiwibmFtZSI6InN0cmluZyBuYW1lIiwiZW1haWwiOiJzdHJpbmcgZW1haWwiLCJnZW5kZXIiOiJzdHJpbmcgZ2VuZGVyIiwiYXZhdGFyIjoic3RyaW5nIGF2YXRhciJ9LCJpc3MiOiJTaGluZVNoYW8iLCJzdWIiOiJ4aWFvc2hhb3FxQGdtYWlsLmNvbSIsImF1ZCI6WyJfQXVkaWVuY2VfIl0sImV4cCI6MTY3MzYyNzU2MywibmJmIjoxNjczNjIwMzYzLCJpYXQiOjE2NzM2MjAzNjMsImp0aSI6IjEifQ.yL8gc7yF-vgVelT-YIwLr4DPOAig6OroKQm-T2picNg"
	user, ok := ParseToken(tokenString)
	if !ok {
		t.Errorf("ParseToken: %v %v", user, ok)
	}
}
