package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func init() {
	InitRedis()
}

var wg sync.WaitGroup

func main() {

	ctx := context.Background()
	err := RClient.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := RClient.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("foo", val)
	wg.Add(1)
	go StartListen(&wg)
	time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {
		RClient.Publish(ctx, "mychannel1", "payload"+strconv.Itoa(i)).Err()
	}
	wg.Wait()
}
