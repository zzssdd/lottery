package model

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"io/ioutil"
	"os"
)

type RedisConf struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

var (
	Client *redis.Client
)

func LoadRedisConf() (Conf *RedisConf) {
	file, err := os.Open("conf/redis_conf.json")
	if err != nil {
		panic(err)
	}
	byte_data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
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
