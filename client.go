package main

import "github.com/redis/go-redis/v9"

var RClient *redis.Client

func InitRedis() {
	RClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
