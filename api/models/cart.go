package models

import "go.mongodb.org/mongo-driver/bson/primitive"

import "errors"

// Cart -> Shopping cart of the client
type Cart struct {
	Shipping int     `json:"shipping"`
	Subtotal int     `json:"subtotal"`
	Total    int     `json:"total"`
	Items    []*Item `json:"items"`
}

// HasProduct -> return true if the card has an item with current productID
func (c *Cart) HasProduct(productID primitive.ObjectID) bool {
	hex := productID.Hex()
	for _, item := range c.Items {
		if item.Product.Hex() == hex {
			return true
		}
	}
	return false
}

// Refresh -> update the cart subtotal and total with the current items
func (c *Cart) Refresh() {
	c.Subtotal = 0
	for _, item := range c.Items {
		c.Subtotal += item.Price * item.Quantity
	}
	c.Total = c.Subtotal + c.Shipping
}

// Remove -> get and item of the cart
func (c *Cart) Remove(product string) error {
	length := len(c.Items)
	for index, item := range c.Items {
		if item.Product.Hex() == product {
			c.Items[index] = c.Items[length-1]
			c.Items = c.Items[:length-1]
			return nil
		}
	}
	return errors.New("product does not exist in cart")
}

type Item struct {
	Product  primitive.ObjectID `json:"product"`
	Name     string             `json:"name"`
	Price    int                `json:"price"`
	Quantity int                `json:"quantity"`
	Size     string             `json:"size"`
}
