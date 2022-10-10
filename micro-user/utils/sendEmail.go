package utils

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

func SendEmail(addr string, code string) error {
	Email := email.NewEmail()
	Email.From = "16546221146@qq.com"
	Email.To = []string{addr}
	Email.Subject = "Yogen抽奖系统用户注册验证"
	Email.Text = []byte(code)
	err := Email.Send("smtp.qq.com:587", smtp.PlainAuth("", "1654622146@qq.com", "afhncxjpsukidgif", "smtp.qq.com"))
	return err
}
