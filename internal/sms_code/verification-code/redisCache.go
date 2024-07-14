package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var (
	ctx     = context.Background()
	redisDb *redis.Client
)

func RedisInit(redisAddr, password string, db int) {
	redisDb = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: password,
		DB:       db,
	})

	pong, err := redisDb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("connect redis_cluster fail: %v", err))
	} else {
		fmt.Printf("connect redis_cluster succ %v\n", pong)
	}

}
