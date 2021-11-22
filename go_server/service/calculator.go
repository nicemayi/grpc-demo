package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"go_server/cache"

	"github.com/go-redis/redis/v8"

	"go_server/service/dto"
)

func Add(ctx context.Context, addRequst *dto.AddRequest) (int, error) {
	log.Printf("In go_server/service/calculator.go, Received: %d, %d, %d", addRequst.A, addRequst.B, addRequst.C)
	return addRequst.A + addRequst.B + addRequst.C, nil
}

func Multiple(ctx context.Context, multipleRequest *dto.MultipleRequest) (int, error) {
	log.Printf("In go_server/service/calculator.go, Received: %v, %d", multipleRequest.A, multipleRequest.B)
	return multipleRequest.A * multipleRequest.B, nil
}

func Fib(ctx context.Context, fibRequest *dto.FibRequest) (int, error) {
	// suppose we have 10 mill-sec delay
	time.Sleep(10 * time.Millisecond)
	numStr := fmt.Sprintf("%d", fibRequest.Number)
	val, err := cache.RedisClient.Get(ctx, numStr).Result()
	if err != nil && err != redis.Nil {
		log.Fatalf("cache error: %v", err)
	}

	var res int
	if err == redis.Nil {
		res, _ := fib(fibRequest.Number)
		err = cache.RedisClient.Set(ctx, numStr, fmt.Sprintf("%d", res), 0).Err()
		if err != nil {
			log.Fatalf("can not ser cache with error: %v", err)
		}
	} else {
		f, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			panic(err)
		}
		res = int(f)
	}

	return res, nil
}

func fib(num int) (int, error) {
	a, b := 0, 1
	if num == 0 {
		return a, nil
	}
	if num == 1 {
		return b, nil
	}

	for i := 0; i < int(num); i++ {
		a, b = b, a+b
	}

	return b, nil
}
