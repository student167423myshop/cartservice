package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Healthz(t *testing.T) {
	// arrange
	initMockRedis()
	defer redisMockServer.Close()
	r := getRouter()
	mockServer := httptest.NewServer(r)

	// act
	resp, _ := http.Get(mockServer.URL + "/healthz")

	// assert
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()
}

func Test_CartAdd_Redis(t *testing.T) {
	// arrange
	initMockRedis()
	defer redisMockServer.Close()
	r := getRouter()
	mockServer := httptest.NewServer(r)

	cart := getCart("0000001", "0000001", 1)
	cartForm := getCartFormSample(cart)

	// act
	resp, _ := http.PostForm(mockServer.URL+"/cart", cartForm)

	// assert
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Status should be 201, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()
}

func Test_CartAddSecondProduct_Redis(t *testing.T) {
	// arrange
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

	// act
	resp, _ := http.PostForm(mockServer.URL+"/cart", cartFormTwo)

	// assert
	bytes, _ := io.ReadAll(resp.Body)
	cart := getCartFromBytes(bytes)

	if clientId != cart.ClientId {
		t.Errorf("ClientId should be %s, got %s",
			clientId, cart.ClientId)
	}

	if firstProductId != cart.CartItems[0].ProductId {
		t.Errorf("CartItem ProductId should be %s, got %s",
			firstProductId, cart.CartItems[0].ProductId)
	}

	if firstProductQuantity != cart.CartItems[0].Quantity {
		t.Errorf("CartItem Quantity should be %d, got %d",
			firstProductQuantity, cart.CartItems[0].Quantity)
	}

	if secondProductId != cart.CartItems[1].ProductId {
		t.Errorf("CartItem ProductId should be %s, got %s",
			secondProductId, cart.CartItems[1].ProductId)
	}

	if secondProductQuantity != cart.CartItems[1].Quantity {
		t.Errorf("CartItem Quantity should be %d, got %d",
			secondProductQuantity, cart.CartItems[1].Quantity)
	}

	lenght := len(cart.CartItems)
	if expectedNumOfProducts != lenght {
		t.Errorf("Cart lenght should be %d, got %d", expectedNumOfProducts, lenght)
	}

	defer resp.Body.Close()
}

func Test_GetCart_Redis(t *testing.T) {
	// arrange
	initMockRedis()
	defer redisMockServer.Close()
	r := getRouter()
	mockServer := httptest.NewServer(r)

	expectedCart := getCart("0000001", "0000001", 1)
	http.PostForm(mockServer.URL+"/cart", getCartFormSample(expectedCart))

	// act
	resp, _ := http.Get(mockServer.URL + "/cart/" + expectedCart.ClientId)

	// assert
	bytes, _ := io.ReadAll(resp.Body)
	cart := getCartFromBytes(bytes)

	if expectedCart.ClientId != cart.ClientId {
		t.Errorf("ClientId should be %s, got %s",
			expectedCart.ClientId, cart.ClientId)
	}

	if expectedCart.CartItems[0].ProductId != cart.CartItems[0].ProductId {
		t.Errorf("CartItem ProductId should be %s, got %s",
			expectedCart.CartItems[0].ProductId, cart.CartItems[0].ProductId)
	}

	if expectedCart.CartItems[0].Quantity != cart.CartItems[0].Quantity {
		t.Errorf("CartItem Quantity should be %d, got %d",
			expectedCart.CartItems[0].Quantity, cart.CartItems[0].Quantity)
	}

	defer resp.Body.Close()
}

func Test_GetEmptyCart_Redis(t *testing.T) {
	// arrange
	initMockRedis()
	defer redisMockServer.Close()
	r := getRouter()
	mockServer := httptest.NewServer(r)

	expectedCart := getEmptyCart("0000001")

	// act
	resp, _ := http.Get(mockServer.URL + "/cart/" + expectedCart.ClientId)

	// assert
	bytes, _ := io.ReadAll(resp.Body)
	cart := getCartFromBytes(bytes)

	if cart.CartItems != nil {
		t.Errorf("CartItems should be nil, got %s", cart.CartItems[0].ProductId)
	}

	defer resp.Body.Close()
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
	cart := getCartFromBytes(bytes)

	if cart.CartItems != nil {
		t.Errorf("CartItems should be nil, got %s", cart.CartItems[0].ProductId)
	}

	defer resp.Body.Close()
}
