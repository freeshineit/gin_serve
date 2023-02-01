package models

import "gorm.io/gorm"

type Todo_Status_Type int

const (
	// 默认
	Todo_Status_Default Todo_Status_Type = 0
	// 选中
	Todo_Status_Checked Todo_Status_Type = 1
)

type Todo struct {
	gorm.Model
	ID      uint             `json:"id" form:"id" gorm:"primary_key:auto_increment"`
	UserID  uint             `json:"user_id" form:"user_id" gorm:"-"`
	Status  Todo_Status_Type `json:"status" form:"status" gorm:"-"`
	Content string           `json:"content" form:"content" binding:"required"`
}
