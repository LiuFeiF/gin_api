package internal

import (
	"fmt"
	"gin_api/pkg/log"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"strings"
)

var RabbitMQ *amqp.Connection

func InitRabbitMQ() {
	name := viper.GetString("amqp.name")
	host := viper.GetString("amqp.host")
	port := viper.GetString("amqp.port")
	username := viper.GetString("amqp.username")
	password := viper.GetString("amqp.password")
	addr := strings.Join([]string{name, "://", username, ":", password, "@", host, ":", port, "/"}, "")
	conn, err := amqp.Dial(addr)
	if err != nil {
		fmt.Println(err)
		log.Error("amqp Dial err:" + err.Error())

	}
	RabbitMQ = conn
	fmt.Println("RabbitMQ 初始化成功")
}
