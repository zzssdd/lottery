package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"lottery/micro-user/proto/user"
	"micro-web/utils"
	"net/http"
	"strconv"
)

func Registry(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	code := c.PostForm("code")
	request := &user.RegistryRequest{
		Email:    email,
		Password: utils.Md5Pw(password),
		Code:     code,
	}
	rep, _ := UserService.Registry(context.TODO(), request)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	request := &user.LoginRequest{
		Email:    email,
		Password: utils.Md5Pw(password),
	}
	rep, _ := UserService.Login(context.TODO(), request)
	if rep.Code != 200 {
		c.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
		return
	}
	token, err := utils.NewToken(email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "生成token错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "登陆成功",
		"token": token,
	})
}

func SendEmail(c *gin.Context) {
	email := c.PostForm("email")
	request := &user.EmailRequest{
		Email: email,
	}
	rep, _ := UserService.SendEmial(context.TODO(), request)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}

func UpdateScore(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	score, _ := strconv.Atoi(c.PostForm("score"))
	request := &user.UpdateRequest{
		Id:    int32(id),
		Score: int32(score),
	}
	rep, _ := UserService.UpdateScore(context.TODO(), request)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}

func GetScore(c *gin.Context) {
	email, exit := c.Get("user")
	if !exit {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "请先登录",
		})
		return
	}
	data := &user.EmailRequest{
		Email: email.(string),
	}
	rep, _ := UserService.GetUserScore(context.TODO(), data)
	if rep.Code != 200 {
		c.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
		"data": rep.Score,
	})
}
