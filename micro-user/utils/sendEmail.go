package utils

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

func SendEmail(addr string, data string) error {

	e := email.NewEmail()
	e.From = "1654622146@qq.com"
	e.To = []string{addr}
	e.Subject = "用户注册验证"
	e.Text = []byte(data)
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "1654622146@qq.com", "afhncxjpsukidgif", "smtp.qq.com"))
	return err
}
