package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Healthz(t *testing.T) {
	// Arrange
	r := getRouter()
	mockServer := httptest.NewServer(r)

	// Act
	resp, _ := http.Get(mockServer.URL + "/healthz")

	// Assert
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func Test_CartAdd_Redis(t *testing.T) {
	// Arrange
	initMockRedis()
	defer redisMockServer.Close()
	r := getRouter()
	mockServer := httptest.NewServer(r)

	cart := getCart("0000001", "0000001", 1)
	cartForm := getCartFormSample(cart)

	// Act
	resp, _ := http.PostForm(mockServer.URL+"/cart", cartForm)

	// Assert
	require.Equal(t, http.StatusCreated, resp.StatusCode)
}

func Test_CartAddSecondProduct_Redis(t *testing.T) {
	// Arrange
	initMockRedis()
	defer redisMockServer.Close()
	r := getRouter()
	mockServer := httptest.NewServer(r)

	clientId := "0000001"
	firstProductId := "0000001"
	firstProductQuantity := 8
	cartFormOne := getCartFormSample(getCart(clientId, firstProductId, firstProductQuantity))
	http.PostForm(mockServer.URL+"/cart", cartFormOne)

	secondProductId := "0000003"
	secondProductQuantity := 2
	expectedNumOfProducts := 2
	cartFormTwo := getCartFormSample(getCart(clientId, secondProductId, secondProductQuantity))

	// Act
	resp, _ := http.PostForm(mockServer.URL+"/cart", cartFormTwo)

	// Assert
	bytes, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	var cart Cart
	_ = json.Unmarshal(bytes, &cart)
	require.Equal(t, clientId, cart.ClientId)
	require.Equal(t, firstProductId, cart.CartItems[0].ProductId)
	require.Equal(t, firstProductQuantity, cart.CartItems[0].Quantity)
	require.Equal(t, secondProductId, cart.CartItems[1].ProductId)
	require.Equal(t, secondProductQuantity, cart.CartItems[1].Quantity)
	require.Equal(t, expectedNumOfProducts, len(cart.CartItems))
}

func Test_GetCart_Redis(t *testing.T) {
	// Arrange
	initMockRedis()
	defer redisMockServer.Close()
	r := getRouter()
	mockServer := httptest.NewServer(r)

	expectedCart := getCart("0000001", "0000001", 1)
	http.PostForm(mockServer.URL+"/cart", getCartFormSample(expectedCart))

	// Act
	resp, _ := http.Get(mockServer.URL + "/cart/" + expectedCart.ClientId)

	// Assert
	bytes, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	var cart Cart
	_ = json.Unmarshal(bytes, &cart)
	require.Equal(t, expectedCart.ClientId, cart.ClientId)
	require.Equal(t, expectedCart.CartItems[0].ProductId, cart.CartItems[0].ProductId)
	require.Equal(t, expectedCart.CartItems[0].Quantity, cart.CartItems[0].Quantity)
}

func Test_GetEmptyCart_Redis(t *testing.T) {
	// Arrange
	initMockRedis()
	defer redisMockServer.Close()
	r := getRouter()
	mockServer := httptest.NewServer(r)

	expectedCart := getEmptyCart("0000001")

	// Act
	resp, _ := http.Get(mockServer.URL + "/cart/" + expectedCart.ClientId)

	// Assert
	bytes, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	var cart Cart
	_ = json.Unmarshal(bytes, &cart)

	require.Equal(t, expectedCart.CartItems, cart.CartItems)
}

func Test_MakeCartEmpty_Redis(t *testing.T) {
	// arrange
	initMockRedis()
	defer redisMockServer.Close()
	r := getRouter()
	mockServer := httptest.NewServer(r)

	expectedCart := getCart("0000001", "0000001", 1)
	http.PostForm(mockServer.URL+"/cart", getCartFormSample(expectedCart))

	// act
	resp, _ := http.Get(mockServer.URL + "/cart/" + expectedCart.ClientId + "/empty")

	// assert
	bytes, _ := io.ReadAll(resp.Body)
	var cart Cart
	_ = json.Unmarshal(bytes, &cart)

	if cart.CartItems != nil {
		t.Errorf("CartItems should be nil, got %s", cart.CartItems[0].ProductId)
	}

	defer resp.Body.Close()
}
