package omipc

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// 监听器结构体
type Listener struct {
	shutdown    chan struct{}
	handler     func(msg string)
	redisClient *redis.Client
}

// 关闭监听器
func (listener Listener) Close() {
	listener.shutdown <- struct{}{}
}

func (listener *Listener) AddHandler(handler func(msg string)) {
	listener.handler = handler
}

// 接受所有发送过来的消息，并执行handler
func (listener *Listener) Listen(channel string) {
	sub := listener.redisClient.Subscribe(context.Background(), channel)
	c := sub.Channel()
	defer sub.Close()
	for {
		select {
		case msg := <-c:
			listener.handler(msg.Payload)
		case <-listener.shutdown:
			return
		}
	}
}
