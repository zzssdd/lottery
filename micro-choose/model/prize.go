package model

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func getRedisName(key string) string {
	return fmt.Sprintf("lottery:prize", key)
}

func Choose() (name string, err error) {
	list_name, err := Client.SMembers(getRedisName("")).Result()
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
	randN := rand.Intn(sum)
	for _, v := range list_name {
		key := getRedisName(v)
		result, _ := Client.HGet(key, "num").Result()
		num, _ := strconv.Atoi(result)
		randN -= num
		if randN < 0 {
			err = Client.HIncrBy(key, "num", -1).Err()
			if err != nil {
				return "", err
			} else {
				return v, nil
			}
		}
	}
	return "", nil
}
