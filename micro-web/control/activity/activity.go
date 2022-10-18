package activity

import (
	"context"
	"github.com/gin-gonic/gin"
	"lottery/micro-activity/proto"
	"net/http"
	"strconv"
	"time"
)

func ActiveAdd(c *gin.Context) {
	name := c.PostForm("name")
	desc := c.PostForm("desc")
	cost, _ := strconv.Atoi(c.PostForm("cost"))
	typ, _ := strconv.Atoi(c.PostForm("type"))
	start := c.PostForm("startTime")
	end := c.PostForm("endTime")
	file, err := c.FormFile("pic")
	user, exit := c.Get("admin")
	if !exit {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请先登陆",
		})
	}
	if err != nil {
		data := &proto.Request{
			Name:      name,
			Desc:      desc,
			User:      user.(string),
			Cost:      int32(cost),
			Type:      int32(typ),
			StartTime: start,
			EndTime:   end,
		}
		rep, _ := ActiveService.ActiveAdd(context.TODO(), data)
		c.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
		return
	}
	unix_int64 := time.Now().Unix()
	unix_str := strconv.FormatInt(unix_int64, 10)
	file_path := "upload/" + unix_str + file.Filename
	err = c.SaveUploadedFile(file, file_path)
	data := &proto.Request{
		Name:      name,
		Desc:      desc,
		User:      user.(string),
		Cost:      int32(cost),
		Type:      int32(typ),
		StartTime: start,
		EndTime:   end,
		Pic:       "http://localhost:8081/" + file_path,
	}
	rep, _ := ActiveService.ActiveAdd(context.TODO(), data)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
		"id":   rep.Id,
	})
}

func ActiveDel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data := &proto.IdRequest{
		Id: int32(id),
	}
	rep, _ := ActiveService.ActiveDel(context.TODO(), data)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}

func ActiveEdit(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	name := c.PostForm("name")
	desc := c.PostForm("desc")
	cost, _ := strconv.Atoi(c.PostForm("cost"))
	typ, _ := strconv.Atoi(c.PostForm("type"))
	start := c.PostForm("startTime")
	end := c.PostForm("endTime")
	file, err := c.FormFile("pic")
	user, exit := c.Get("admin")
	if !exit {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请先登陆",
		})
	}
	if err != nil {
		data := &proto.EditRequest{
			Id:        int32(id),
			Name:      name,
			Desc:      desc,
			User:      user.(string),
			Cost:      int32(cost),
			Type:      int32(typ),
			StartTime: start,
			EndTime:   end,
		}
		rep, _ := ActiveService.ActiveEdit(context.TODO(), data)
		c.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
		return
	}
	unix_int64 := time.Now().Unix()
	unix_str := strconv.FormatInt(unix_int64, 10)
	file_path := "upload/" + unix_str + file.Filename
	err = c.SaveUploadedFile(file, file_path)
	data := &proto.EditRequest{
		Id:        int32(id),
		Name:      name,
		Desc:      desc,
		User:      user.(string),
		Cost:      int32(cost),
		Type:      int32(typ),
		StartTime: start,
		EndTime:   end,
		Pic:       "http://localhost:8081/" + file_path,
	}
	rep, _ := ActiveService.ActiveEdit(context.TODO(), data)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}

func ActiveList(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "6"))
	data := &proto.ListRequest{
		PageNum:  int32(pageNum),
		PageSize: int32(pageSize),
	}
	rep, _ := ActiveService.ActiveList(context.TODO(), data)
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
		"data":  rep.Acts,
		"count": rep.Count,
	})
}

func ActiveInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data := &proto.IdRequest{
		Id: int32(id),
	}
	rep, _ := ActiveService.ActiveInfo(context.TODO(), data)
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
		"data": rep.Act,
	})
}

func ActivePrizes(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	data := &proto.IdRequest{
		Id: int32(id),
	}
	rep, _ := ActiveService.ActivePrizes(context.TODO(), data)
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
		"data": rep.Prizes,
	})
}

func ActiveAddPrize(c *gin.Context) {
	name := c.PostForm("name")
	id, _ := strconv.Atoi(c.PostForm("id"))
	data := &proto.PrizeAddRequest{
		Name: name,
		Id:   int32(id),
	}
	rep, _ := ActiveService.ActiveAddPrize(context.TODO(), data)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}
