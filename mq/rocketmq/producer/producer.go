package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"strconv"
)

func main() {
	addr,err := primitive.NewNamesrvAddr("127.0.0.1:9876")
	if err != nil {
		panic(err)
	}
	topic := "Develop"
	p,err := rocketmq.NewProducer(
		producer.WithGroupName("my_service"),
		producer.WithNameServer(addr),
		producer.WithCreateTopicKey(topic),
		producer.WithRetry(1))
	if err != nil {
		panic(err)
	}

	err = p.Start()
	if err != nil {
		panic(err)
	}

	// 发送异步消息
	res,err := p.SendSync(context.Background(),primitive.NewMessage(topic,[]byte("send sync message")))
	if err != nil {
		fmt.Printf("send sync message error:%s\n",err)
	} else {
		fmt.Printf("send sync message success. result=%s\n",res.String())
	}

	// 发送消息后回调
	err = p.SendAsync(context.Background(), func(ctx context.Context, result *primitive.SendResult, err error) {
		if err != nil {
			fmt.Printf("receive message error:%v\n",err)
		} else {
			fmt.Printf("send message success. result=%s\n",result.String())
		}
	},primitive.NewMessage(topic,[]byte("send async message")))
	if err != nil {
		fmt.Printf("send async message error:%s\n",err)
	}

	// 批量发送消息
	var msgs []*primitive.Message
	for i := 0; i < 5; i++ {
		msgs = append(msgs, primitive.NewMessage(topic,[]byte("batch send message. num:"+strconv.Itoa(i))))
	}
	res,err = p.SendSync(context.Background(),msgs...)
	if err != nil {
		fmt.Printf("batch send sync message error:%s\n",err)
	} else {
		fmt.Printf("batch send sync message success. result=%s\n",res.String())
	}

	// 发送延迟消息
	msg := primitive.NewMessage(topic,[]byte("delay send message"))
	msg.WithDelayTimeLevel(3)
	res,err = p.SendSync(context.Background(),msg)
	if err != nil {
		fmt.Printf("delay send sync message error:%s\n",err)
	} else {
		fmt.Printf("delay send sync message success. result=%s\n",res.String())
	}

	// 发送带有tag的消息
	msg1 := primitive.NewMessage(topic,[]byte("send tag message"))
	msg1.WithTag("tagA")
	res,err = p.SendSync(context.Background(),msg1)
	if err != nil {
		fmt.Printf("send tag sync message error:%s\n",err)
	} else {
		fmt.Printf("send tag sync message success. result=%s\n",res.String())
	}

	err = p.Shutdown()
	if err != nil {
		panic(err)
	}
}