package choose

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"micro-web/rabbitmq"
	"micro-web/redis"
	"net/http"
	"strconv"
)

func Choose(c *gin.Context) {
	email, exist := c.Get("user")
	id, _ := strconv.Atoi(c.PostForm("id"))
	if !exist {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "请先登陆",
		})
		return
	}
	map_data := map[string]interface{}{
		"email": email,
		"id":    id,
	}
	byte_data, _ := json.Marshal(map_data)
	msg := string(byte_data)
	err := rabbitmq.Publish(msg)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "抽奖失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "抽奖中，请稍后",
	})
}

func ChooseResult(c *gin.Context) {
	email, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "token不存在",
		})
		return
	}
	result, err := redis.Get(email.(string))
	if err != nil {
		fmt.Println(err)
		fmt.Println("**")
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "暂未获取到抽奖结果",
		})
		return
	}
	err = redis.Del(email.(string))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "暂未获取到抽奖结果",
		})
		return
	}
	var rep map[string]interface{}
	json.Unmarshal([]byte(result), &rep)
	fmt.Println(rep)
	c.JSON(http.StatusOK, gin.H{
		"code": rep["code"],
		"msg":  rep["msg"],
		"data": rep["prize"],
	})
}
