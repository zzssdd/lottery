package middleware

import (
	"github.com/gin-gonic/gin"
	"micro-web/utils"
	"net/http"
	"strings"
)

func AdminJwt(c *gin.Context) {
	auth_header := c.Request.Header.Get("Authorization")
	if auth_header == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请先登录",
		})
		c.Abort()
		return
	}
	auths := strings.Split(auth_header, " ")
	if len(auths) < 2 {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请先登录",
		})
		c.Abort()
		return
	}
	user, err := utils.AuthToken(auths[1])
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "该用户没有权限",
		})
		c.Abort()
		return
	}
	c.Set("admin", user.Username)
	c.Next()
}
func UserJwt(c *gin.Context) {
	auth_header := c.Request.Header.Get("Authorization")
	if auth_header == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请先登录",
		})
		c.Abort()
		return
	}
	auths := strings.Split(auth_header, " ")
	if len(auths) < 2 {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请先登录",
		})
		c.Abort()
		return
	}
	user, err := utils.AuthToken(auths[1])
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "该用户没有权限",
		})
		c.Abort()
		return
	}
	c.Set("user", user.Username)
	c.Next()

}
