package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Cart -> Shopping cart of the client
type Cart struct {
	Shipping int
	SubTotal int
	Total    int
	Items    []*Item
}

// Compute -> update the cart subtotal and total with the current items
func (c *Cart) Compute() {
	c.SubTotal = 0
	for _, item := range c.Items {
		c.SubTotal += item.Price
	}
	c.Total = c.SubTotal + c.Shipping

}

type Item struct {
	Product  primitive.ObjectID
	Name     string
	Price    int
	Quantity int
	Size     string
}
