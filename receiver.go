package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func StartListen(wg *sync.WaitGroup) {
	ctx := context.Background()
	pubsub := RClient.Subscribe(ctx, "mychannel1")
	print("LISTENING")
	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
		time.Sleep(2 * time.Second)
	}
}
