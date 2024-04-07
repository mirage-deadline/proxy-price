package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	// test client to check provider messages
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	sub := rdb.Subscribe(ctx, "crypto-prices")
	defer sub.Close()
	for msg := range sub.Channel() {
		fmt.Println(msg.Payload)
	}
}
