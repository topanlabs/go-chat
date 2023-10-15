package main

import (
	"context"
	"fmt"
	"log"
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

func deliverMessages(id string) {
	ctx := context.Background()

	pubsub := RClient.Subscribe(ctx, id)
	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
		gStore.findAndDeliver(msg.Channel, string(msg.Payload))
	}
}
func (s *Store) findAndDeliver(userID string, content string) {
	m := Message{
		Content: content,
	}
	for _, u := range s.Users {
		if u.ID == userID {
			if err := u.conn.WriteJSON(m); err != nil {
				log.Printf("error on message delivery e: %s\n", err)
			} else {
				log.Printf("user %s found, message sent\n", userID)
			}
			return
		}
	}
	log.Printf("user %s not found at our store\n", userID)
}
