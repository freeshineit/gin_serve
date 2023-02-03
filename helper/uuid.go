package helper

import "github.com/google/uuid"

func GenTodoUuId() string {
	uuid := uuid.New()
	return uuid.String()
}
