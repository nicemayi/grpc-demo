package cache

import (
	"fmt"
	"log"
	"os"

	redis "github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var RedisClient *redis.Client

func init() {
	REDIS_URL := os.Getenv("REDIS_URL")
	REDIS_PORT := os.Getenv("REDIS_PORT")
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", REDIS_URL, REDIS_PORT),
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
