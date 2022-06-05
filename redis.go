package main

import (
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
)

var redisClient *redis.Client = nil
var redisMockServer *miniredis.Miniredis

func getRedis() *redis.Client {
	if redisClient == nil {
		initNewRedis()
	}

	return redisClient
}

func initNewRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func initMockRedis() {
	redisMockServer, _ = miniredis.Run()
	redisClient = redis.NewClient(&redis.Options{
		Addr: redisMockServer.Addr(),
	})
}
