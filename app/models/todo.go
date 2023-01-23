package models

const (
	Todo_STATUS_DEFAULT int = 0
	Todo_STATUS_DELETE  int = 1
)

type Todo struct {
	Id      int    `json:"id" form:"id"`
	Status  int    `json:"status" form:"status"`
	Content string `json:"content" form:"content" binding:"required"`
}
