package models

type User struct {
	Name   string `json:"name" form:"name" binding:"required"`
	Email  string `json:"email" form:"email" binding:"required"`
	Gender string `json:"gender" form: "gender" binding:"required"`
	Avatar string `json:"avatar" form: "avatar" binding:"required"`
}
