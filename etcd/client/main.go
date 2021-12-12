package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"time"
)

var EtcdClient *clientv3.Client

func init() {
	var err error
	EtcdClient, err = clientv3.New(clientv3.Config{Endpoints: []string{"localhost:2379"},
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to etcd success!")

}

func main() {

	DistributedLock()
}

func basicOption() {
	key := "k1"
	value := "v2"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	go Watch(ctx, key)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	Put(ctx, key, value)
	cancel()

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	Get(ctx, key)
	cancel()

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	Lease(key, value)
	cancel()

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	Del(ctx, key)
	cancel()

	time.Sleep(time.Second * 10)
}

func Del(ctx context.Context, key string) {
	delResponse, err := EtcdClient.Delete(ctx, key)
	if err != nil {
		fmt.Printf("etcd del to etcd failed, err: %#v", err)
	}
	fmt.Println("delete count: ", delResponse.Deleted)
}

func Get(ctx context.Context, key string) {
	resp, err := EtcdClient.Get(ctx, key)

	if err != nil {
		fmt.Printf("etcd get to etcd failed, err: %#v", err)
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("etcd get value: %s:%s\n", ev.Key, ev.Value)
	}

}

func Put(ctx context.Context, key, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := EtcdClient.Put(ctx, key, value)
	cancel()

	if err != nil {
		fmt.Printf("etcd put to etcd failed, err: %#v", err)
	}
	fmt.Println("Put success!")
}

func Watch(ctx context.Context, key string) {
	fmt.Printf("watch start :%s\n", key)
	watchCh := EtcdClient.Watch(ctx, key)
	for wresp := range watchCh {
		for _, env := range wresp.Events {
			// 获取被修改的key
			fmt.Printf("Watch type:%s key:%s value: %s\n", env.Type, env.Kv.Key, env.Kv.Value)
		}
	}
}

// 五秒租约
func Lease(key, value string) {
	fmt.Printf("Lease start ,%s:%s\n", key, value)

	resp, err := EtcdClient.Grant(context.TODO(), 5)
	if err != nil {
		panic(err)
	}
	_, err = EtcdClient.Put(context.TODO(), key, value, clientv3.WithLease(resp.ID))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Lease end ,%s:%s\n", key, value)
}

func DistributedLock() {
	sessionOne, err := concurrency.NewSession(EtcdClient)

	if err != nil {
		panic(err)
	}
	defer sessionOne.Close()
	myLockOne := concurrency.NewMutex(sessionOne, "/my-lock/")

	sessionTwo, err := concurrency.NewSession(EtcdClient)

	if err != nil {
		panic(err)
	}
	defer sessionTwo.Close()

	myLockTwo := concurrency.NewMutex(sessionTwo, "/my-lock/")

	// 会话1获取锁
	if err := myLockOne.Lock(context.TODO()); err != nil {
		panic(err)
	}
	fmt.Println("acquired lock for myLockOne")

	myLocked := make(chan struct{})

	go func() {
		defer close(myLocked)
		// 等待直到会话sessionOne释放了/my-lock/的锁
		if err := myLockTwo.Lock(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := myLockOne.Unlock(context.TODO()); err != nil {
		panic(err)
	}
	fmt.Println("released lock for sessionOne")

	<-myLocked
	fmt.Println("acquired lock for myLockTwo")

}
