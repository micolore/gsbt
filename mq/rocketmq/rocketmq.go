package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

var rocketAdmin admin.Admin
var err error

func init() {
	nameSrvAddr := []string{"192.168.1.170:9876"}
	rocketAdmin, err = admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(nameSrvAddr)))
	if err != nil {
		panic(err)
	}
}
func main() {
	topic := "Develop"
	brokerAddr := "192.168.1.170:10911"

	if err != nil {
		panic(err)
	}

	// 创建topic
	err = rocketAdmin.CreateTopic(
		context.Background(),
		admin.WithTopicCreate(topic),
		admin.WithBrokerAddrCreate(brokerAddr))
	if err != nil {
		fmt.Println("Create topic error:", err)
	}

}

func RemoveTopic(topic string) {
	// 删除topic
	err = rocketAdmin.DeleteTopic(
		context.Background(),
		admin.WithTopicDelete(topic),
		//admin.WithBrokerAddrDelete(brokerAddr),
		//admin.WithNameSrvAddr(nameSrvAddr),
	)

	err = rocketAdmin.Close()
	if err != nil {
		fmt.Println("Shutdown admin error:", err)
	}
}
