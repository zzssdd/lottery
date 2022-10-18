package rabbitmq

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

type RabbitMqConf struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Vhost    string `json:"vhost"`
}

func LoadRabbitConf() (Conf *RabbitMqConf) {
	file, err := os.Open("conf/rabbitmq_conf.json")
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
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%s/%s",
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
	Ch, _ = Conn.Channel()
	Ch.QueueDeclare("order_queue", true, false, false, false, nil)
	Ch.ExchangeDeclare("order_exchange", amqp.ExchangeDirect, true, false, false, false, nil)
	Ch.QueueBind("order_queue", "", "order_exchange", false, nil)
}

func Publish(msg string) error {
	err := Ch.Publish("order_exchange", "", false, false, amqp.Publishing{
		ContentType:  "text/plain",
		Body:         []byte(msg),
		DeliveryMode: amqp.Persistent,
	})
	return err
}

func Close() {
	Conn.Close()
	Ch.Close()
}
