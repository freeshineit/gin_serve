package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID     uint   `json: "id" form:"name" gorm:"primary_key:auto_increment"`
	Name   string `json:"name" form:"name" binding:"required" gorm:"type:varchar(255)"`
	Email  string `json:"email" form:"email" binding:"required" gorm:"unique_index;type:varchar(255)"`
	Age    string `json:"age" form:"age" binding:"required" gorm:"type:int"`
	Gender string `json:"gender" form: "gender" binding:"required" gorm:"type:int"`
	Avatar string `json:"avatar" form: "avatar" binding:"required" gorm:"type:varchar(255)"`
}

// 用户登录
type UserLogin struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// 用户登录
type UserRegister struct {
	User
	Password        string `json:"password" form:"password" binding:"required" gorm:"->;<-;not null" validate:"min=6,max=30"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required" validate:"min=6,max=30,eqfield=Password"`
}

type LoginRecord struct {
	gorm.Model
	ID     uint   `gorm:"primary_key:auto_increment"`
	UserID uint64 `gorm:"not null"`
	Token  string `gorm:"-"`
}
