package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Picture struct {
	Reference string      `json:"reference"`
	Src       string      `json:"src"`
	Data      primitive.M `json:"data"`
}
