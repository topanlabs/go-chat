package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	gStore *Store
)

func init() {
	InitRedis()
	gStore = &Store{
		Users: make([]*User, 0, 1),
	}
}

func main() {

	ctx := context.Background()
	err := RClient.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/ws", wsHandler)
	r.Run()
}

func wsHandler(c *gin.Context) {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("upgrader error %s\n" + err.Error())
		return
	}
	u := gStore.newUser(conn)
	log.Printf("user %s joined\n", u.ID)
	go deliverMessages(u.ID)
	time.Sleep(1 * time.Second)
	ctx := context.Background()
	RClient.Publish(ctx, u.ID, "Your id is: "+u.ID)
	for {
		var m Message
		if err := u.conn.ReadJSON(&m); err != nil {
			log.Printf("error on ws. message %s\n", err)
		}

		RClient.Publish(ctx, m.DeliveryID, string(m.Content))
	}
}
