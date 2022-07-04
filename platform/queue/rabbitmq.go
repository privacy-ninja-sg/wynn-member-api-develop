package queue

import (
	"github.com/streadway/amqp"
	"wynn-member-api/pkg/configs"
	"wynn-member-api/pkg/utils"
)

func RabbitmqConnection() *amqp.Connection {
	url := configs.RabbitmqConfig()
	conn, err := amqp.Dial(url)
	utils.FailOnErr(err)
	return conn
}
