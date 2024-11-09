package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	omipc "github.com/stormi-li/omipc"
)

func main() {
	go process1()
	process2()
}

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

func process1() {
	l := omipc.NewClient(&redis.Options{Addr: redisAddr, Password: password}).NewLock("lock")
	for i := 0; i < 1000; i++ {
		l.Lock()
		fmt.Println("process1占有锁")
		time.Sleep(1 * time.Second)
		l.Unlock()
		fmt.Println("process1释放锁")
		time.Sleep(10 * time.Millisecond)
	}
}

func process2() {
	l := omipc.NewClient(&redis.Options{Addr: redisAddr, Password: password}).NewLock("lock")
	for i := 0; i < 1000; i++ {
		l.Lock()
		fmt.Println("process2占有锁")
		time.Sleep(1 * time.Second)
		l.Unlock()
		fmt.Println("process2释放锁")
		time.Sleep(10 * time.Millisecond)
	}
}
