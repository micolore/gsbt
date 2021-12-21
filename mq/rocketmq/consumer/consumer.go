package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"os"
	"time"
)

// 在v2.1.0-rc5.0不支持，会在下一个版本中支持
func PullConsumer() {
	topic := "Develop"

	// 消费者主动拉取消息
	// not
	c1, err := rocketmq.NewPullConsumer(
		consumer.WithGroupName("my_service"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"my-rq-mq:9876"})))
	if err != nil {
		panic(err)
	}
	err = c1.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	queue := primitive.MessageQueue{
		Topic:      topic,
		BrokerName: "broker-a", // 使用broker的名称
		QueueId:    0,
	}

	err = c1.Shutdown()
	if err != nil {
		fmt.Println("Shutdown Pull Consumer error: ", err)
	}

	offset := int64(0)
	for {
		resp, err := c1.PullFrom(context.Background(), queue, offset, 10)
		if err != nil {
			if err == rocketmq.ErrRequestTimeout {
				fmt.Printf("timeout\n")
				time.Sleep(time.Second)
				continue
			}
			fmt.Printf("unexpected error: %v\n", err)
			return
		}
		if resp.Status == primitive.PullFound {
			fmt.Printf("pull message success. nextOffset: %d\n", resp.NextBeginOffset)
			for _, ext := range resp.GetMessageExts() {
				fmt.Printf("pull msg: %s\n", ext)
			}
		}
		offset = resp.NextBeginOffset
	}
}

func PushConsumer() {
	topic := "Develop"

	// 消息主动推送给消费者
	c2, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("my_service"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"192.168.150.70:9876"})),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset), // 选择消费时间(首次/当前/根据时间)
		consumer.WithConsumerModel(consumer.BroadCasting)) // 消费模式(集群消费:消费完其他人不能再读取/广播消费：所有人都能读)
	if err != nil {
		panic(err)
	}

	err = c2.Subscribe(
		topic, consumer.MessageSelector{
			Type:       consumer.TAG,
			Expression: "*", // 可以 TagA || TagB
		},
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			orderlyCtx, _ := primitive.GetOrderlyCtx(ctx)
			fmt.Printf("orderly context: %v\n", orderlyCtx)
			for i := range msgs {
				fmt.Printf("Subscribe callback: %v\n", msgs[i])
			}
			return consumer.ConsumeSuccess, nil
		})
	if err != nil {
		fmt.Printf("Subscribe error:%s\n", err)
	}

	err = c2.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	time.Sleep(time.Minute)
	err = c2.Shutdown()
	if err != nil {
		fmt.Println("Shutdown Consumer error: ", err)
	}
}
