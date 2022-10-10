package control

import (
	"context"
	"github.com/patrickmn/go-cache"
	"micro-user/model"
	"micro-user/proto/user"
	"micro-user/utils"
	"time"
)

type User struct {
}

var code = cache.New(60*time.Second, 10*time.Second)

func (u *User) Registry(ctx context.Context, in *user.RegistryRequest, out *user.Response) error {
	code, ok := code.Get(in.Email)
	if !ok {
		out.Code = 500
		out.Msg = "未获取验证码"
		return nil
	}
	if code != in.Code {
		out.Code = 500
		out.Msg = "验证码输入不正确，请重新输入"
		return nil
	}
	data := &model.User{
		Email:    in.Email,
		Password: in.Password,
	}
	err := model.UserRegistry(data)
	if err != nil {
		out.Code = 500
		out.Msg = "注册用户失败"
		return err
	}
	out.Code = 200
	out.Msg = "用户注册成功"
	return nil
}
func (u *User) Login(ctx context.Context, in *user.LoginRequest, out *user.Response) error {
	data := &model.User{
		Email:    in.Email,
		Password: in.Password,
	}
	is_ok := model.UserLogin(data)
	if !is_ok {
		out.Code = 500
		out.Msg = "账号密码输入错误"
		return nil
	}
	out.Code = 200
	out.Msg = "登陆成功"
	return nil
}
func (u *User) SendEmial(ctx context.Context, in *user.EmailRequest, out *user.Response) error {
	email := in.Email
	is_ok := model.UserExist(email)
	if !is_ok {
		out.Code = 500
		out.Msg = "该邮箱已被注册"
		return nil
	}
	num := utils.RandNum(6)
	code.Set(email, code, cache.DefaultExpiration)
	err := utils.SendEmail(email, num)
	if err != nil {
		out.Code = 500
		out.Msg = "发送验证码失败，请稍后再试"
		return err
	}
	out.Code = 200
	out.Msg = "发送验证码成功"
	return nil
}
