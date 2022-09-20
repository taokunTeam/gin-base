package bootstrap

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/taokunTeam/gin-base/global"
	"go.uber.org/zap"
)

func InitializeRabbitMq() *amqp.Connection {
	mq := fmt.Sprintf("amqp://%s:%s@%s:%d/", global.App.Config.RabbitMq.Name, global.App.Config.RabbitMq.Password, global.App.Config.RabbitMq.Host, global.App.Config.RabbitMq.Port)
	conn, err := amqp.Dial(mq)
	if err != nil {
		global.App.Log.Error("Failed to connect to RabbitMQ, err:", zap.Any("err", err))
		return nil
	}
	//defer conn.Close()
	global.App.Log.Info("connect  RabbitMQ success!")
	return conn
}
