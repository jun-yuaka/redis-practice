package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	// Set a key
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		fmt.Println("Set error", err)
	} else {
		fmt.Println("Key set successfully.")
	}

	for {
		val, err := rdb.Get(ctx, "key").Result()
		if err != nil {
			fmt.Println("Get error:", err)
		} else {
			fmt.Println("key", val)
		}

		// ここで5秒間スリープする
		time.Sleep(5 * time.Second)
	}
}
