package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Guest struct {
	ID       primitive.ObjectID `bson:"_id"`
	Cart     *Cart
	Shipping *Shipping
}
