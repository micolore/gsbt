package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {

	//delKey()
	delSingDel()

}

func delSingDel() {
	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "r-d9j7jjde90xhbhzop6.redis.ap-southeast-5.rds.aliyuncs.com:6379",
		Password: "Fb19Otu2doxHpoa",
		DB:       0,
	})
	fmt.Println(rdb)
	val := rdb.Get(ctx, "wsn:bus:remainBytes:9")
	fmt.Println(val)
	val2 := rdb.Del(ctx, "wsn:bus:remainBytes:9")
	fmt.Println(val2)
	fmt.Println("del success!")
}
func delKey() {
	var ctx = context.Background()

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"192.168.11.199:8000", "192.168.11.199:8001", "192.168.11.199:8002", "192.168.11.199:8003", "192.168.11.199:8004", "192.168.11.199:8005"},
		//To route commands by latency or randomly, enable one of the following.
		//RouteByLatency: true,
		//RouteRandomly: true,
	})

	for i := 0; i < 100000; i++ {
		key := fmt.Sprintf("lopdeals:goods:attr:%d", i)
		rdb.Del(ctx, key)
		fmt.Println(key)
	}
}
