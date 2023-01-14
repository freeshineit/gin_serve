package models

type User struct {
	Id     string `json: "id" form:"name"`
	Name   string `json:"name" form:"name" binding:"required"`
	Email  string `json:"email" form:"email" binding:"required"`
	Gender string `json:"gender" form: "gender" binding:"required"`
	Avatar string `json:"avatar" form: "avatar" binding:"required"`
}

// 用户登录
type UserLogin struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// 用户登录
type UserRegister struct {
	Name            string `json:"name" form:"name" binding:"required"`
	Email           string `json:"email" form:"email" binding:"required"`
	Gender          string `json:"gender" form: "gender" binding:"required"`
	Avatar          string `json:"avatar" form: "avatar" binding:"required"`
	Password        string `json:"password" form:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required"`
}
