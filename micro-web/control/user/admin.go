package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"lottery/micro-user/proto/admin"
	"micro-web/utils"
	"net/http"
	"strconv"
)

func AdminLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	request := &admin.LoginRequest{
		Username: username,
		Password: utils.Md5Pw(password),
	}
	rep, _ := AdminService.AdminLogin(context.TODO(), request)
	if rep.Code != 200 {
		c.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
		return
	}
	token, err := utils.NewToken(username)
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

func UserList(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	request := &admin.UsersRequest{
		PageNum:  int32(pageNum),
		PageSize: int32(pageSize),
	}
	rep, _ := AdminService.UserList(context.TODO(), request)
	if rep.Code != 200 {
		c.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  rep.Code,
		"msg":   rep.Msg,
		"data":  rep.Users,
		"count": rep.Count,
	})
}

func UserDel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	request := &admin.DelRequest{
		Id: int32(id),
	}
	rep, _ := AdminService.UserDel(context.TODO(), request)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}
