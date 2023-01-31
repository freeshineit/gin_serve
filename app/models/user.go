package models

type User struct {
	ID     uint   `json: "id" form:"name" gorm:"primary_key:auto_increment"`
	Name   string `json:"name" form:"name" binding:"required" gorm:"type:varchar(255)"`
	Email  string `json:"email" form:"email" binding:"required" gorm:"uniqueIndex;type:varchar(255)"`
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
	ID              uint   `json: "id" form:"name" gorm:"primary_key:auto_increment"`
	Name            string `json:"name" form:"name" binding:"required"`
	Email           string `json:"email" form:"email" binding:"required"`
	Gender          string `json:"gender" form: "gender" binding:"required"`
	Avatar          string `json:"avatar" form: "avatar" binding:"required"`
	Password        string `json:"password" form:"password" binding:"required" gorm:"->;<-;not null"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required" `
}

type LoginRecord struct {
	ID     uint   `gorm:"primary_key:auto_increment"`
	UserID uint64 `gorm:"not null"`
	Token  string `gorm:"-"`
}
