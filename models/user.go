package models

type User struct {
	Name   string `json:"name" form:"name`
	Email  string `json:"email" form:"email"`
	Gender string `json:"gender" form: "gender"`
	Avatar string `json:"avatar" form: "avatar"`
}
