package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getEmptyCart(t *testing.T) {
	// Arrange
	expectedClientId := "00001"

	// Act
	cart := getEmptyCart(expectedClientId)

	// Assert
	require.Equal(t, expectedClientId, cart.ClientId)
	require.Empty(t, cart.CartItems)
}

func Test_getCart(t *testing.T) {
	// Arrange
	expectedClientId := "00001"
	expectedProductId := "00002"
	expectedQuantity := 2

	// Act
	cart := getCart(
		expectedClientId,
		expectedProductId,
		expectedQuantity,
	)

	// Assert
	require.Equal(t, expectedClientId, cart.ClientId)
	require.Equal(t, expectedProductId, cart.CartItems[0].ProductId)
	require.Equal(t, expectedQuantity, cart.CartItems[0].Quantity)
}

func Test_getCartFormSample(t *testing.T) {
	// Arrange
	expectedClientId := "00001"
	expectedProductId := "00002"
	quantity := 2
	expectedQuantity := "2"
	cart := getCart(
		expectedClientId,
		expectedProductId,
		quantity,
	)

	// Act
	cartSample := getCartFormSample(cart)

	// Assert
	require.Equal(t, expectedClientId, cartSample.Get("userId"))
	require.Equal(t, expectedProductId, cartSample.Get("productId"))
	require.Equal(t, expectedQuantity, cartSample.Get("quantity"))
}

func Test_addCartItemToCart(t *testing.T) {
	// Arrange
	expectedClientId := "0000001"
	expectedCartItemOne := CartItem{"00001", 1}
	expectedCartItemTwo := CartItem{"00002", 4}
	var expectedCartItems []CartItem
	expectedCartItems = append(expectedCartItems, expectedCartItemTwo)
	expectedCartItems = append(expectedCartItems, expectedCartItemOne)
	expectedCart := Cart{
		expectedClientId,
		expectedCartItems,
	}
	cart := getEmptyCart(expectedClientId)

	// Act
	cart = addCartItemToCart(cart, CartItem{"00002", 1})
	cart = addCartItemToCart(cart, CartItem{"00001", 1})
	cart = addCartItemToCart(cart, CartItem{"00002", 3})

	// Assert
	require.Equal(t, expectedCart, cart)
}

func Test_addCartItemToCart_MultipleTimes(t *testing.T) {
	// Arrange
	clientId := "0000001"
	firstProductId := "0000001"
	firstProductQuantity := 8
	secondProductId := "0000003"
	secondProductQuantity := 2
	expectedNumOfProducts := 2
	cart := getEmptyCart(clientId)

	// Act
	cart = addCartItemToCart(cart, CartItem{firstProductId, 2})
	cart = addCartItemToCart(cart, CartItem{firstProductId, 4})
	cart = addCartItemToCart(cart, CartItem{secondProductId, 2})
	cart = addCartItemToCart(cart, CartItem{firstProductId, 2})

	// Assert
	require.Equal(t, clientId, cart.ClientId)
	require.Equal(t, firstProductId, cart.CartItems[0].ProductId)
	require.Equal(t, firstProductQuantity, cart.CartItems[0].Quantity)
	require.Equal(t, secondProductId, cart.CartItems[1].ProductId)
	require.Equal(t, secondProductQuantity, cart.CartItems[1].Quantity)
	require.Equal(t, expectedNumOfProducts, len(cart.CartItems))
}

func Test_addCartItemVariablesToCart_AddAndSub(t *testing.T) {
	// Arrange
	clientId := "0000001"
	firstProductId := "0000001"
	firstProductQuantity := 4
	secondProductId := "0000003"
	expectedNumOfProducts := 1
	cart := getEmptyCart(clientId)

	// Act
	cart = addCartItemToCart(cart, CartItem{firstProductId, 2})
	cart = addCartItemToCart(cart, CartItem{firstProductId, 4})
	cart = addCartItemToCart(cart, CartItem{secondProductId, 2})
	cart = addCartItemToCart(cart, CartItem{secondProductId, -2})
	cart = addCartItemToCart(cart, CartItem{firstProductId, -2})

	// Assert
	require.Equal(t, clientId, cart.ClientId)
	require.Equal(t, firstProductId, cart.CartItems[0].ProductId)
	require.Equal(t, firstProductQuantity, cart.CartItems[0].Quantity)
	require.Equal(t, expectedNumOfProducts, len(cart.CartItems))
}
