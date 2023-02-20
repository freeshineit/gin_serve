package dto

type CaptchaDTO struct {
	CaptchaID string `json:"captcha_id" form:"captcha_id"` //验证码Id
	ImageURL  string `json:"image_url" form:"image_url"`   //验证码图片url
}
