package main

import (
	"testing"
)

func Test_CartAddMultipleProducts(t *testing.T) {
	// arrange
	clientId := "0000001"
	firstProductId := "0000001"
	firstProductQuantity := 8
	secondProductId := "0000003"
	secondProductQuantity := 2
	expectedNumOfProducts := 2
	cart := getEmptyCart(clientId)

	// act
	cart = addCartItemVariablesToCart(cart, firstProductId, 2)
	cart = addCartItemVariablesToCart(cart, firstProductId, 4)
	cart = addCartItemVariablesToCart(cart, secondProductId, 2)
	cart = addCartItemVariablesToCart(cart, firstProductId, 2)

	// assert
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
		t.Errorf("Cart lenght should be %d, got %d",
			expectedNumOfProducts, lenght)
	}
}

func Test_CartAddAndSubMultipleProducts(t *testing.T) {
	// arrange
	clientId := "0000001"
	firstProductId := "0000001"
	firstProductQuantity := 4
	secondProductId := "0000003"
	expectedNumOfProducts := 1
	cart := getEmptyCart(clientId)

	// act
	cart = addCartItemVariablesToCart(cart, firstProductId, 2)
	cart = addCartItemVariablesToCart(cart, firstProductId, 4)
	cart = addCartItemVariablesToCart(cart, secondProductId, 2)
	cart = addCartItemVariablesToCart(cart, secondProductId, -2)
	cart = addCartItemVariablesToCart(cart, firstProductId, -2)

	// assert
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
	lenght := len(cart.CartItems)
	if expectedNumOfProducts != lenght {
		t.Errorf("Cart lenght should be %d, got %d",
			expectedNumOfProducts, lenght)
	}
}
