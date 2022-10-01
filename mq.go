package main

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main1() {
	p, err := rocketmq.NewProducer(producer.WithNameServer([]string{"127.0.0.1:9876"}))
	if err != nil {
		panic("生成 producer 失败")
	}
	if err = p.Start(); err != nil {
		panic("启动 producer 失败")
	}
	res, err := p.SendSync(context.Background(), primitive.NewMessage("RocketMQ", []byte("this is RocketMQ")))
	if err != nil {
		fmt.Printf("发送失败: %s\n", err)
	} else {
		fmt.Printf("发送成功: %s\n", res.String())
	}
	if err = p.Shutdown(); err != nil {
		panic("关闭 producer 失败")
	}
}
