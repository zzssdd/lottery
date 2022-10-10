package model

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"io/ioutil"
	"os"
	"time"
)

var (
	Client *redis.Client
)

type RedisConf struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func LoadRedisConf() (Conf *RedisConf) {
	file, err := os.Open("conf/redis_conf.json")
	if err != nil {
		panic(err)
	}
	byte_data, err := ioutil.ReadAll(file)
	err = json.Unmarshal(byte_data, &Conf)
	if err != nil {
		panic(err)
	}
	return
}

func init() {
	conf := LoadRedisConf()
	Client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", conf.Host, conf.Port),
	})
}

func Set(key string, value string) {
	Client.Set(key, value, 1*time.Hour)
}
