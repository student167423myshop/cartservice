package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getHealthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func GetCartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientId := vars["userId"]
	cart := getCartFromRedis(clientId)
	w.Header().Set("Content-Type", "application/json")
	w.Write(getWriteCart(cart))
}

func AddItemToCartHandler(w http.ResponseWriter, r *http.Request) {
	cart := getCartFromForm(r)
	for _, cartItem := range cart.CartItems {
		cart = addCartItemToRedis(cart.ClientId, cartItem)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(getWriteCart(cart))
}

func EmptyCartHandler(w http.ResponseWriter, r *http.Request) {
	clientId := mux.Vars(r)["userId"]
	cart := getEmptyCart(clientId)
	cart = replaceCartInRedis(cart)
	w.WriteHeader(http.StatusOK)
	w.Write(getWriteCart(cart))
}

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

func replaceCartInRedis(cart Cart) Cart {
	client := getRedis()
	jsonData := getWriteCart(cart)
	_ = client.Set(cart.ClientId, jsonData, 0).Err()
	return getCartFromRedis(cart.ClientId)
}

func addCartItemToRedis(clientId string, cartItem CartItem) Cart {
	cart := getCartFromRedis(clientId)
	cart = addCartItemToCart(cart, cartItem)
	json, _ := json.Marshal(cart)
	client := getRedis()
	_ = client.Set(clientId, json, 0).Err()
	return getCartFromRedis(clientId)
}

func getCartFromForm(r *http.Request) Cart {
	var clientId = r.FormValue("userId")
	var productId = r.FormValue("productId")
	var quantityStr = r.FormValue("quantity")
	quantity, _ := strconv.Atoi(quantityStr)
	var cart = getCart(clientId, productId, quantity)
	return cart
}
