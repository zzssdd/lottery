package model

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func getRedisName(index int) string {
	key := strconv.Itoa(index)
	return fmt.Sprintf("lottery:prize:%s", key)
}

func Choose(id int) (name string, err error) {
	list_name, err := Client.SMembers(getRedisName(id)).Result()
	if err != nil {
		return "", err
	}
	rand.Seed(time.Now().Unix())
	sum := 0
	for _, v := range list_name {
		result, _ := Client.HGet(v, "num").Result()
		num, _ := strconv.Atoi(result)
		sum += num
	}
	randN := rand.Intn(sum + 1)
	for _, v := range list_name {
		result, _ := Client.HGet(v, "num").Result()
		num, _ := strconv.Atoi(result)
		randN -= num
		if randN <= 0 {
			err = Client.HIncrBy(v, "num", -1).Err()
			if err != nil {
				return "", err
			} else {
				return v, nil
			}
		}
	}
	return "", nil
}
