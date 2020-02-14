package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Cart -> Shopping cart of the client
type Cart struct {
	Shipping int         `json:"shipping"`
	Subtotal int         `json:"subtotal"`
	Total    int         `json:"total"`
	Items    []*CartItem `json:"items"`
	Executed bool        `json:"-"`
}

// Refresh -> update the cart subtotal and total with the current items
func (c *Cart) Refresh() {
	c.Subtotal = 0
	for _, item := range c.Items {
		c.Subtotal += item.Price * item.Quantity
	}
	c.Total = c.Subtotal + c.Shipping
}

// HasItem -> return true if the card has an item with current productID
func (c *Cart) HasItem(id primitive.ObjectID) bool {
	hex := id.Hex()
	for _, item := range c.Items {
		if item.ID.Hex() == hex {
			return true
		}
	}
	return false
}

// AddItem -> add an item to the cart
func (c *Cart) AddItem(item *CartItem) error {
	if c.HasItem(item.ID) {
		return fmt.Errorf("product '%s' is already added to the cart", item.ID.Hex())
	}
	c.Items = append(c.Items, item)
	c.Refresh()
	return nil
}

// RemoveItem -> remove an item of the cart
func (c *Cart) RemoveItem(id primitive.ObjectID) bool {
	length := len(c.Items)
	for index, item := range c.Items {
		if item.ID.Hex() == id.Hex() {
			c.Items[index] = c.Items[length-1]
			c.Items = c.Items[:length-1]
			c.Refresh()
			return true
		}
	}
	return false
}
