package main

import (
	"sync"

	"github.com/google/uuid"

	"github.com/gorilla/websocket"
)

type User struct {
	ID   string
	conn *websocket.Conn
}
type Store struct {
	Users []*User
	sync.Mutex
}

type Message struct {
	DeliveryID string `json:"id"`
	Content    string `json:"content"`
}

func (s *Store) newUser(conn *websocket.Conn) *User {
	u := &User{
		ID:   uuid.New().String(),
		conn: conn,
	}
	s.Lock()
	defer s.Unlock()
	s.Users = append(s.Users, u)
	return u
}
