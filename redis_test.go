package main

import (
	"testing"
)

func TestRedisConnection(t *testing.T) {
	initMockRedis()
	defer redisMockServer.Close()
	client := getRedis()
	_, err := client.Ping().Result()
	if err != nil {
		t.Fatal(err)
	}
}
