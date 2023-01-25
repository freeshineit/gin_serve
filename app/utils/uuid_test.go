package utils

import (
	"fmt"
	"testing"
)

// Example creating a token using a custom claims type.  The StandardClaim is embedded
// in the custom type to allow for easy encoding, parsing and validation of standard claims.
func TestUuid(t *testing.T) {
	// tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjoiIiwibmFtZSI6IlNoaW5lU2hhbyIsImVtYWlsIjoieGlhb3NoYW9xcUBnbWFpbC5jb20iLCJnZW5kZXIiOiJNYW4iLCJhdmF0YXIiOiIvdXBsb2FkL2F2YXRhci5wbmcifSwiaXNzIjoiU2hpbmVTaGFvIiwic3ViIjoieGlhb3NoYW9xcUBnbWFpbC5jb20iLCJhdWQiOlsiX0F1ZGllbmNlXyJdLCJleHAiOjE2NzM3MDk2MzEsIm5iZiI6MTY3MzcwMjQzMSwiaWF0IjoxNjczNzAyNDMxLCJqdGkiOiIxIn0.pXTxq-KsYgWTUtKkE8SjMqpEqEiQNSj-JFt3UEFHM7A"

	uuidStr := GenTodoUuId()

	fmt.Println(uuidStr)

	t.Errorf("uuid: %s", uuidStr)
}
