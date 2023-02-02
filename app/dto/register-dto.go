package dto

type RegisterDTO struct {
	Name            string `json:"name" form:"name" binding:"required"`
	Email           string `json:"email" form:"email" binding:"required"`
	Age             string `json:"age" form:"age" binding:"required"`
	Gender          string `json:"gender" form: "gender" binding:"required"`
	Avatar          string `json:"avatar" form: "avatar" binding:"required"`
	Password        string `json:"password" form:"password" binding:"required,min=6,max=30"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required,min=6,max=30,eqfield=Password"`
}
