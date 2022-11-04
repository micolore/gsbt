package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"strconv"
	"testing"
	"time"
)

var tunnel = make(chan string, 1) // 取决于函数业务的参数类型

func Test66(t *testing.T) {
	go IncomingZombie()

	chairPool, _ := ants.NewPoolWithFunc(3, ExecuteZombie) // 声明有几把电椅
	defer chairPool.Release()

	for {
		select {
		case a := <-tunnel:
			go chairPool.Invoke(a) //将参数传输给执行函数
			fmt.Println("当前执行的协程数： ", chairPool.Running())
		}
	}
}

// 处决僵尸
func ExecuteZombie(i interface{}) {
	fmt.Printf("正在处决僵尸 %s 号，还有5秒钟....\n", i.(string))
	time.Sleep(1 * time.Second)
	fmt.Printf(":) %s 玩完了，下一个\n-----------------\n", i.(string))
}

// 僵尸不断进来
func IncomingZombie() {
	for i := 0; i < 4; i++ {
		tunnel <- strconv.Itoa(i)
	}
}
