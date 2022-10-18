package redis

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"io/ioutil"
	"os"
)

var RedisClient *redis.Client

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
	RedisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", conf.Host, conf.Port),
	})
}

func Close() {
	RedisClient.Close()
}

func Get(email string) (string, error) {
	return RedisClient.Get(email).Result()
}

func Del(email string) error {
	return RedisClient.Del(email).Err()
}
