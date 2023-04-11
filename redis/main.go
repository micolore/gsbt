package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	//delKey()
	//delSingDel()
	zAdd()
}

func delSingDel() {
	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	fmt.Println(rdb)
	val2 := rdb.Del(ctx, "key")
	fmt.Println(val2)
	fmt.Println("del success!")

}
func zAdd() {
	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	for i := 0; i < 10; i++ {
		err := rdb.ZAdd(ctx, "myset", redis.Z{
			Score:  float64(i),
			Member: fmt.Sprintf("member%d", i),
		}).Err()
		if err != nil {
			panic(err)
		}
	}
	keys, _, err := rdb.ZScan(ctx, "myset", 3, "", 4).Result()
	if err != nil {
		panic(err)
	}
	println(keys[12])
}
func delClusterKey() {
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
