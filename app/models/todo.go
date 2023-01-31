package models

type Todo_Status_Type int

const (
	// 默认
	Todo_Status_Default Todo_Status_Type = 0
	// 选中
	Todo_Status_Checked Todo_Status_Type = 1
)

type Todo struct {
	Id      string           `json:"id" form:"id"`
	Status  Todo_Status_Type `json:"status" form:"status"`
	Content string           `json:"content" form:"content" binding:"required"`
}
