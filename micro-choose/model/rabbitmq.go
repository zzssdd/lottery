package model

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"io/ioutil"
	"os"
)

var (
	Conn *amqp.Connection
	Ch   *amqp.Channel
	err  error
)

type RabbitConf struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Vhost    string `json:"vhost"`
}

func LoadRabbitConf() (Conf *RabbitConf) {
	file, err := os.Open("conf/rabbit_conf.json")
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
	conf := LoadRabbitConf()
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%s//%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Vhost,
	)
	Conn, err = amqp.Dial(dsn)
	if err != nil {
		panic(err)
	}
	Ch, err = Conn.Channel()
	if err != nil {
		panic(err)
	}
	Ch.Qos(1, 0, false)
}
