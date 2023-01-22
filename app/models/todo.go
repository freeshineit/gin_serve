package models

type Todo struct {
	Id      string `json:"id" form:"id"`
	Status  string `json:"status" form:"status"`
	Content string `json:"content" form:"content"`
}
