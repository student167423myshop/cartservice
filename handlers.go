package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getHealthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getCartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientId := vars["userId"]
	cart := getCartFromRedis(clientId)
	jsonCart, err := json.Marshal(cart)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonCart)
	}
}

func addItemToCartHandler(w http.ResponseWriter, r *http.Request) {
	cart := getCartFromForm(r)
	for _, cartItem := range cart.CartItems {
		cart = addCartItemToRedis(cart.ClientId, cartItem)
	}
	jsonCart, err := json.Marshal(cart)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonCart)
	}
}

func emptyCartHandler(w http.ResponseWriter, r *http.Request) {
	clientId := mux.Vars(r)["userId"]
	cart := getEmptyCart(clientId)
	cart, err := replaceCartInRedis(cart)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	jsonCart, err := json.Marshal(cart)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonCart)
	}
}
