package models

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCartAdd(t *testing.T) {
	cart := &Cart{}
	cartItem := &CartItem{
		ID:       primitive.NewObjectID(),
		Price:    1000,
		Quantity: 2,
	}
	cart.AddItem(cartItem)

	if len(cart.Items) != 1 {
		t.Errorf("expected cart items length to be 1, received %d", len(cart.Items))
	}

	if err := cart.AddItem(cartItem); err == nil {
		t.Error("Should return error if try to add an item that is already added")
	}

	if cart.Total != 2000 {
		t.Errorf("cart total should be 2000, received: %d", cart.Total)
	}
}

func TestCartRefresh(t *testing.T) {
	testcases := []struct {
		cart  *Cart
		total int
	}{
		{
			cart: &Cart{
				Shipping: 200,
				Items: []*CartItem{
					&CartItem{ID: primitive.NewObjectID(), Price: 1000, Quantity: 2},
					&CartItem{ID: primitive.NewObjectID(), Price: 500, Quantity: 2},
				},
			},
			total: 3200,
		}, {
			cart: &Cart{
				Shipping: 100,
				Items: []*CartItem{
					&CartItem{ID: primitive.NewObjectID(), Price: 200, Quantity: 3},
					&CartItem{ID: primitive.NewObjectID(), Price: 800, Quantity: 2},
				},
			},
			total: 2300,
		},
	}
	for _, testcase := range testcases {
		testcase.cart.Refresh()
		if testcase.total != testcase.cart.Total {
			t.Errorf("expected total to be %d, received: %d", testcase.total, testcase.cart.Total)
		}
	}
}

func TestCartRemove(t *testing.T) {
	id := primitive.NewObjectID()
	cartItem := &CartItem{
		ID:       id,
		Price:    2000,
		Quantity: 2,
	}
	cart := &Cart{
		Items: []*CartItem{cartItem},
	}
	deleted := cart.RemoveItem(id)

	if !deleted {
		t.Errorf("expected cart.RemoveItem to return true, received %v", deleted)
	}

	if len(cart.Items) != 0 {
		t.Errorf("expected cart items length to be zero, received: %d", len(cart.Items))
	}

	if cart.RemoveItem(id) {
		t.Error("expected cart.RemoveItem to return false, since there's no items in list")
	}
}

func TestHasItem(t *testing.T) {
	id := primitive.NewObjectID()
	cart := &Cart{
		Items: []*CartItem{&CartItem{ID: id}},
	}
	if !cart.HasItem(id) {
		t.Error("expect cart.HasItem to return true on existen item")
	}
	if cart.HasItem(primitive.NewObjectID()) {
		t.Error("expected cart.HasItem to return false on unexistent item")
	}
}
