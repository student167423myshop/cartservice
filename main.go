package main

import (
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func getApp() *fiber.App {
	app := fiber.New()
	app.Get("/api/v1/cart/:userId", getCart)
	app.Get("/api/v1/cart/:userId/add/:productId/:quantity", addItem)
	app.Get("/api/v1/cart/:userId/empty", emptyCart)
	return app
}

func getRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return client
}

func main() {
	app := getApp()
	app.Listen(":7070")
}
