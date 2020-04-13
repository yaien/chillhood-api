package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Picture struct {
	ID   primitive.ObjectID `bson:"_id" json:"id"`
	Src  string             `json:"src"`
	Data primitive.M        `json:"data"`
}
