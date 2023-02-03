package model

type User struct {
	Model
	ID       uint   `json:"id" form:"id" gorm:"primary_key:auto_increment"`
	Name     string `json:"name" form:"name" binding:"required" gorm:"type:varchar(255)"`
	Email    string `json:"email" form:"email" binding:"required" gorm:"uniqueIndex;type:varchar(255)"`
	Age      string `json:"age" form:"age" binding:"required" gorm:"type:int"`
	Gender   string `json:"gender" form:"gender" binding:"required" gorm:"type:int"`
	Avatar   string `json:"avatar" form:"avatar" binding:"required" gorm:"type:varchar(255)"`
	Password string `json:"-" gorm:"->;<-;not null"`
	// Token    string `json:"token,omitempty" from: "token,omitempty" gorm:"-"`
}
