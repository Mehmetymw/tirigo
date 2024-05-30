package redis

import (
	"context"

	"github.com/go-redis/redis"
)

var Client *redis.Client
var Ctx = context.Background()

func Init(addr string) {
	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	_, err := Client.Ping().Result()
	if err != nil {
		panic(err)
	}
}
