package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// CartItem -> an item of the cart
type CartItem struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `json:"name"`
	Price       int                `json:"price"`
	Quantity    int                `json:"quantity"`
	Size        string             `json:"size"`
	Picture     *Picture           `json:"picture"`
	Description string             `json:"description"`
}
