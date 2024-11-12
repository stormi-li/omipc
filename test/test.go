package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omipc"
)

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

func main() {
	c := omipc.NewClient(&redis.Options{Addr: redisAddr, Password: password})
	c.Listen("channel", func(message string) bool {
		fmt.Println(message)
		return false
	})
}
func Pub() {
	redisClient := redis.NewClient(&redis.Options{Addr: redisAddr, Password: password})
	redisClient.Publish(context.Background(), "channel", "hello")
}
func Sub() {
	redisClient := redis.NewClient(&redis.Options{Addr: redisAddr, Password: password})
	sub := redisClient.Subscribe(context.Background(), "channel")
	c := sub.Channel()
	time.Sleep(3 * time.Second)
	fmt.Println("Sub")
	fmt.Println(<-c)
}
