package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/cart", addItemToCartHandler).Methods(http.MethodPost)
	r.HandleFunc("/cart/{userId}", getCartHandler).Methods(http.MethodGet)
	r.HandleFunc("/cart/{userId}/empty", emptyCartHandler).Methods(http.MethodGet)
	r.HandleFunc("/healthz", getHealthz).Methods(http.MethodGet)
	return r
}

func main() {
	r := getRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    ":7070",
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}
