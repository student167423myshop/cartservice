package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// STRUCTS
type Cart struct {
	ClientId  string     `json:"ClientId"`
	CartItems []CartItem `json:"CartItems"`
}

type CartItem struct {
	ProductId string `json:"ProductId"`
	Quantity  int    `json:"Quantity"`
}

// API
func getCartFromRedis(clientId string) Cart {
	client := getRedis()
	var cart Cart
	jsonData, err := client.Get(clientId).Result()
	if err != nil {
		cart = getEmptyCart(clientId)
	} else {
		_ = json.Unmarshal([]byte(jsonData), &cart)
	}
	return cart
}

func replaceCartInRedis(cart Cart) (Cart, error) {
	var replacedCart Cart
	client := getRedis()
	jsonCart, err := json.Marshal(cart)
	if err != nil {

		return replacedCart, err
	}
	_ = client.Set(cart.ClientId, jsonCart, 0).Err()
	replacedCart = getCartFromRedis(cart.ClientId)
	return replacedCart, nil
}

func addCartItemToRedis(clientId string, cartItem CartItem) Cart {
	cart := getCartFromRedis(clientId)
	cart = addCartItemToCart(cart, cartItem)
	json, _ := json.Marshal(cart)
	client := getRedis()
	_ = client.Set(clientId, json, 0).Err()
	return getCartFromRedis(clientId)
}

// INNER FUNCTIONS
func getEmptyCart(clientId string) Cart {
	var cartItems []CartItem
	cart := Cart{clientId, cartItems}
	return cart
}

func getCart(clientId string, productId string, quantity int) Cart {
	cartItem := CartItem{productId, quantity}
	cartItems := []CartItem{cartItem}
	cart := Cart{clientId, cartItems}
	return cart
}

func getCartFormSample(cart Cart) url.Values {
	form := url.Values{}
	form.Add("userId", cart.ClientId)
	form.Add("productId", cart.CartItems[0].ProductId)
	form.Add("quantity", strconv.Itoa(cart.CartItems[0].Quantity))
	return form
}

func addCartItemToCart(cart Cart, newCartItem CartItem) Cart {
	found := false
	if newCartItem.Quantity > 0 {
		for index, cartItem := range cart.CartItems {
			if cartItem.ProductId == newCartItem.ProductId {
				cartItem.Quantity += newCartItem.Quantity
				cart.CartItems[index].Quantity = cartItem.Quantity
				found = true
			}
		}
		if !found {
			cart.CartItems = append(cart.CartItems, newCartItem)
		}
	} else {
		cart = subCartItemFromCart(cart, newCartItem)
	}
	return cart
}

func subCartItemFromCart(cart Cart, newCartItem CartItem) Cart {
	if newCartItem.Quantity < 0 {
		for index, cartItem := range cart.CartItems {
			if cartItem.ProductId == newCartItem.ProductId {
				cartItem.Quantity += newCartItem.Quantity
				if cartItem.Quantity <= 0 {
					cart.CartItems = removeCartItem(cart.CartItems, index)
				} else {
					cart.CartItems[index].Quantity = cartItem.Quantity
				}
			}
		}
	}
	return cart
}

func removeCartItem(cartItems []CartItem, index int) []CartItem {
	return append(cartItems[:index], cartItems[index+1:]...)
}

func getCartFromForm(r *http.Request) Cart {
	var clientId = r.FormValue("userId")
	var productId = r.FormValue("productId")
	var quantityStr = r.FormValue("quantity")
	quantity, _ := strconv.Atoi(quantityStr)
	var cart = getCart(clientId, productId, quantity)
	return cart
}
