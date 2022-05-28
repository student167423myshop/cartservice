package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Cart struct {
	ClientId  string
	CartItems []CartItem
}

type CartItem struct {
	ProductId string
	Quantity  int
}

func getCart(c *fiber.Ctx) error {
	var clientId = c.Params("userId")
	client := getRedis()

	cartJsonFromRedis, err := client.Get(clientId).Result()
	if err != nil {
		fmt.Println(err)
	}

	var cartFromRedis Cart
	json.Unmarshal([]byte(cartJsonFromRedis), &cartFromRedis)

	return c.JSON(cartFromRedis)
}

func addItem(c *fiber.Ctx) error {
	var clientId = c.Params("userId")
	var productId = c.Params("productId")
	var quantityStr = c.Params("quantity")

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		fmt.Println(err)
	}

	var cartItem = CartItem{
		ProductId: productId,
		Quantity:  quantity,
	}

	client := getRedis()

	cartJsonFromRedis, err := client.Get(clientId).Result()
	if err != nil {
		fmt.Println(err)
	}

	var cartFromRedis Cart
	json.Unmarshal([]byte(cartJsonFromRedis), &cartFromRedis)

	cartFromRedis.CartItems = append(cartFromRedis.CartItems, cartItem)

	newCartJson, err := json.Marshal(cartFromRedis)
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set(clientId, newCartJson, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func emptyCart(c *fiber.Ctx) error {
	var clientId = c.Params("userId")

	client := getRedis()

	cartJsonFromRedis, err := client.Get(clientId).Result()
	if err != nil {
		fmt.Println(err)
	}

	var cartFromRedis Cart
	json.Unmarshal([]byte(cartJsonFromRedis), &cartFromRedis)

	cartFromRedis.CartItems = nil

	newCartJson, err := json.Marshal(cartFromRedis)
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set(clientId, newCartJson, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
