package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Guest struct {
	ID   primitive.ObjectID `bson:"_id" json:"id"`
	Cart *Cart              `json:"cart,omitempty"`
}
