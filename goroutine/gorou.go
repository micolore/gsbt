package main

import (
	"fmt"
	databases "moppo.com/gsbt/mysql"
	"sync"
	"time"
)
// channel生产者
func ChannelProducer(logChan chan InsertLog, log InsertLog, wait *sync.WaitGroup) {
	logChan <- log
	fmt.Println("product data：", log)
	wait.Done()
}
// channel消费者
func ChannelConsumer(logChan chan InsertLog, wait *sync.WaitGroup) {
	log := <-logChan
	fmt.Println("consumer data：", log)
	SaveInsertLog(log)
	wait.Done()
}

func main() {
	var wg sync.WaitGroup
	logChan := make(chan InsertLog, 100)

	for i := 0; i < 10; i++ {
		fmt.Printf("Producer %d\n", i)
		il := InsertLog{
			Type:        3,
			Description: "我的描述",
			Goroutine:   i,
			CreateAt:    time.Now(),
		}
		go ChannelProducer(logChan, il, &wg)
		wg.Add(1)
	}
	for j := 0; j < 10; j++ {
		fmt.Printf("Consumer %d\n", j)
		go ChannelConsumer(logChan, &wg)
		wg.Add(1)
	}
	wg.Wait()
}

// InsertLog
type InsertLog struct {
	Id          int       `json:"id" gorm:"column:id;type:bigint;primary_key"`
	Type        int       `json:"type" gorm:"column:type;type:bigint;"`                              //类型
	Goroutine   int       `json:"goroutine" gorm:"column:goroutine;type:bigint;"`                    //所属协程
	TaskStart   int       `json:"taskStart" gorm:"column:task_start;type:bigint;"`                   //任务开始
	TaskEnd     int       `json:"taskEnd" gorm:"column:task_end;type:bigint;"`                       //任务结束
	CreateBy    int       `json:"taskEnd" gorm:"column:create_by;type:bigint;"`                      //创建人
	CreateAt    time.Time `json:"taskEnd" gorm:"column:create_at;type:timestamp without time zone;"` //创建时间
	Description string    `json:"description" gorm:"column:description;type:character varying;"`     //描述信息
}

func SaveInsertLog(insertLog InsertLog) {
	databases.DB.Save(&insertLog)
}
