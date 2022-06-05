package main

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type Cart struct {
	ClientId  string
	CartItems []CartItem
}

type CartItem struct {
	ProductId string
	Quantity  int
}

func getEmptyCart(clientId string) Cart {
	var cartItems []CartItem
	cart := Cart{clientId, cartItems}
	return cart
}

func getWriteCart(cart Cart) []byte {
	jsonData, _ := json.Marshal(cart)
	return jsonData
}

func getCartFromBytes(bytes []byte) Cart {
	var cart Cart
	_ = json.Unmarshal(bytes, &cart)
	return cart
}

func getCart(clientId string, productId string, quantity int) Cart {
	cartItem := CartItem{productId, quantity}
	cartItems := []CartItem{cartItem}
	cart := Cart{clientId, cartItems}
	return cart
}

func getCartItem(productId string, quantity int) CartItem {
	cartItem := CartItem{productId, quantity}
	return cartItem
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

func addCartItemVariablesToCart(cart Cart, productId string, quantity int) Cart {
	cartItem := getCartItem(productId, quantity)
	cart = addCartItemToCart(cart, cartItem)
	return cart
}

func removeCartItem(cartItems []CartItem, index int) []CartItem {
	return append(cartItems[:index], cartItems[index+1:]...)
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
