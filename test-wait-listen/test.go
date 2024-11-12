package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	omipc "github.com/stormi-li/omipc"
)

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

func main() {
	listen()
	wait()
}

func wait() {
	omipc := omipc.NewClient(&redis.Options{Addr: redisAddr, Password: password})
	for i := 0; i < 10000; i++ {
		fmt.Println(omipc.Wait("notifyall", 1*time.Second))
		time.Sleep(1000 * time.Millisecond)
	}
}

func listen() {
	omipc := omipc.NewClient(&redis.Options{Addr: redisAddr, Password: password})
	omipc.Listen("notifyall", func(message string) bool {
		fmt.Println(message)
		return true
	})
}
