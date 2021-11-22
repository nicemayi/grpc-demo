package cache

import (
	"log"

	redis "github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	res := RedisClient.Ping(context.Background())
	if res.Err() != nil {
		log.Panicf("Can not start redis!")
		pong, err := res.Result()
		if err != nil || pong != "pong" {
			log.Panicf("Can not start redis!")
		}
	}
}
