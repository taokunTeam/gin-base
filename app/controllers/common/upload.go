package common

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/taokunTeam/gin-base/app/common/request"
	"github.com/taokunTeam/gin-base/app/common/response"
	"github.com/taokunTeam/gin-base/app/services"
	"github.com/taokunTeam/gin-base/global"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func ImageUpload(c *gin.Context) {
	var form request.ImageUpload
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	outPut, err := services.MediaService.SaveImage(form)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, outPut)
}

func Test(c *gin.Context) {
	ch, err := global.App.RabbitMQ.Channel()
	global.App.Log.Error("Failed to open a channel:", zap.Any("err", err))
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	global.App.Log.Error("Failed to declare a queue", zap.Any("err", err))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	global.App.Log.Error("Failed to publish a message", zap.Any("err", err))
	global.App.Log.Error("Failed to publish a message", zap.Any("err", body))
	log.Printf(" [x] Sent %s\n", body)
}
