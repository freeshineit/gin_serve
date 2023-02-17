package helper

import (
	"fmt"
	"gin_serve/app/dto"
	"gin_serve/config"

	"gopkg.in/gomail.v2"
)

type Email struct {
	To      []string
	Subject string
	Attach  []string
}

func SendEmail(to []string, subject string, bodyHtml string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "serve_service@163.com")
	// m.SetHeader("To", "842383741@qq.com", "767027642@qq.com")
	m.SetHeader("To", to...)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	// m.SetHeader("Subject", "Hello! hahaha")
	m.SetHeader("Subject", subject)
	// m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	m.SetBody("text/html", bodyHtml)
	// m.Attach("/home/Alex/lolcat.jpg")  // 附件

	d := gomail.NewDialer("smtp.163.com", 25, "serve_service@163.com", config.Conf.Email.Password)

	// Send the email to `to`
	return d.DialAndSend(m)
}

func SendActiveEmail(user *dto.UserDTO) error {
	token, err := GenerateEmailToken(user.ID, user.Email)
	if err != nil {
		return err
	}

	return SendEmail([]string{user.Email}, "Verify Active Email", fmt.Sprintf("<div><div>Click Active Email</div><a href='http://localhost:8080/verify_email/%s'>http://localhost:8080/verify_email/%s</a></div>", token, token))
}
