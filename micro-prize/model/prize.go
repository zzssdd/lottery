package model

import (
	"fmt"
	"strconv"
)

type Prize struct {
	Name       string
	Num        int
	Pic        string
	CreateTime string
}

func getRedisName(key string) string {
	return fmt.Sprintf("lottery:prize", key)
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
	name := getRedisName(prize.Name)
	err := Client.SAdd(getRedisName(""), name).Err()
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	Client.HSet(name, "num", prize.Num)
	Client.HSet(name, "pic", prize.Pic)
	Client.HSet(name, "createTime", prize.CreateTime)
	return nil
}

func PrizeDel(name string) error {
	name = getRedisName(name)
	err := Client.SRem(getRedisName(""), name).Err()
	if err != nil {
		return err
	}
	err = Client.HDel(name, "probability").Err()
	if err != nil {
		return err
	}
	Client.HDel(name, "num")
	Client.HDel(name, "pic")
	Client.HDel(name, "createTime")
	return nil
}

func PrizeEdit(name string, prize *Prize) error {
	name = getRedisName(name)
	err := Client.SRem(getRedisName(""), name).Err()
	if err != nil {
		return err
	}
	err = Client.SAdd(getRedisName(""), prize.Name).Err()
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	Client.HSet(name, "num", prize.Num)
	Client.HSet(name, "pic", prize.Pic)
	Client.HSet(name, "createTime", prize.CreateTime)
	return nil
}

func PrizeInfo(name string) (prize *Prize, err error) {
	name = getRedisName(name)
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
