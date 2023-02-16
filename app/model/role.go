package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID          uint64 `json:"id" form:"id" gorm:"primary_key:auto_increment"`
	Name        string `json:"name" form:"name" binding:"required" gorm:"type:varchar(255);not null"`
	Description string `json:"description" form:"description" binding:"required" gorm:"type:text"`
	Status      *uint  `json:"status" form:"status" gorm:"not null;default:0"` // 收否激活
}
