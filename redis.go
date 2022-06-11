package main

import (
	"os"

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
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	password := os.Getenv("REDIS_PASS")
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
}

func initMockRedis() {
	redisMockServer, _ = miniredis.Run()
	redisClient = redis.NewClient(&redis.Options{
		Addr: redisMockServer.Addr(),
	})
}
