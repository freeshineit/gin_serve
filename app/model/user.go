package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint64 `json:"id" form:"id" gorm:"primary_key:auto_increment"`
	Name     string `json:"name" form:"name" binding:"required" gorm:"type:varchar(255)"`
	Email    string `json:"email" form:"email" binding:"required" gorm:"uniqueIndex;type:varchar(255)"`
	Age      uint   `json:"age" form:"age" binding:"required" gorm:"type:int"`
	Gender   string `json:"gender" form:"gender" binding:"required" gorm:"type:varchar(40)"`
	Avatar   string `json:"avatar" form:"avatar" binding:"required" gorm:"type:varchar(255)"`
	Password string `json:"-" gorm:"->;<-;not null"`
}
