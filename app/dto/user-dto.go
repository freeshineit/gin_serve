package dto

type UserDTO struct {
	Model
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Age      uint   `json:"age" form:"age" binding:"required,gte=1,lte=200"`
	Gender   string `json:"gender" form:"gender" binding:"required"`
	Avatar   string `json:"avatar" form:"avatar" binding:"required"`
	IsActive *uint  `json:"is_active" form:"is_active"` // 收否激活 0 未激活 1 已激活
	RoleId   *uint  `json:"role_id" form:"role_id"`     // 角色 1 默认角色
}

// UserRegisterDTO is used by when user register
type UserRegisterDTO struct {
	Name            string `json:"name" form:"name" binding:"required"`
	Email           string `json:"email" form:"email" binding:"required,email"`
	Age             uint   `json:"age" form:"age" binding:"required,gte=1,lte=200"`
	Gender          string `json:"gender" form:"gender" binding:"required"`
	Avatar          string `json:"avatar" form:"avatar" binding:"required"`
	Password        string `json:"password" form:"password" binding:"required,min=6,max=30"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required,min=6,max=30,eqfield=Password"`
}

// UserUpdateDTO is used by when update profile
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Age      uint   `json:"age" form:"age" binding:"required,gte=1,lte=200"`
	Gender   string `json:"gender" form:"gender" binding:"required"`
	Avatar   string `json:"avatar" form:"avatar" binding:"required"`
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required,min=6,max=40"`
}

// UserLoginDTO is used by when user login
type UserLoginDTO struct {
	Email     string `json:"email" form:"email" binding:"required,email"`
	Password  string `json:"password" form:"password" binding:"required,min=6,max=40"`
	Code      string `json:"code" form:"code" binding:"required"`
	CaptchaID string `json:"captcha_id" form:"captcha_id" binding:"required"`
}
