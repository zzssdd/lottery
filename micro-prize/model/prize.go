package model

import (
	"fmt"
	"strconv"
	"time"
)

type Prize struct {
	Name       string
	Num        int
	Pic        string
	CreateTime string
}

func getRedisName(key string) string {
	return fmt.Sprintf("lottery:prize:%s", key)
}

func PrizeList() (prizes []*Prize, err error) {
	list_name, err := Client.SMembers(getRedisName("")).Result()
	for _, v := range list_name {
		value, _ := Client.HGetAll(v).Result()
		num, _ := strconv.Atoi(value["num"])
		prize := &Prize{
			Name:       v,
			Num:        num,
			Pic:        value["pic"],
			CreateTime: value["createTime"],
		}
		prizes = append(prizes, prize)
	}
	return
}

func PrizeAdd(prize *Prize) error {
	name := prize.Name
	err := Client.SAdd(getRedisName(""), name).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err != nil {
		fmt.Println(err)
		return err
	}
	Client.HSet(name, "num", prize.Num)
	Client.HSet(name, "pic", prize.Pic)
	Client.HSet(name, "createTime", prize.CreateTime)
	return nil
}

func PrizeDel(name string) error {
	err := Client.SRem(getRedisName(""), name).Err()
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	Client.HDel(name, "num")
	Client.HDel(name, "pic")
	Client.HDel(name, "createTime")
	return nil
}

func PrizeEdit(name string, prize *Prize) error {
	err := Client.SRem(getRedisName(""), name).Err()
	if err != nil {
		return err
	}
	Client.HDel(name, "num")
	Client.HDel(name, "pic")
	createTime, err := Client.HGet(name, "createTime").Result()
	if err != nil {
		createTime = time.Now().Format("2006-01-02 15:04:05")
	}
	Client.HDel(name, "createTime")
	err = Client.SAdd(getRedisName(""), prize.Name).Err()
	if err != nil {
		return err
	}
	Client.HSet(prize.Name, "num", prize.Num)
	if err != nil {
		return err
	}
	Client.HSet(prize.Name, "pic", prize.Pic)
	Client.HSet(prize.Name, "createTime", createTime)
	return nil
}

func PrizeInfo(name string) (prize *Prize, err error) {
	value, _ := Client.HGetAll(name).Result()
	num, _ := strconv.Atoi(value["num"])
	prize = &Prize{
		Name:       name,
		Num:        num,
		Pic:        value["pic"],
		CreateTime: value["createTime"],
	}
	return
}
