package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Province struct {
	ID   primitive.ObjectID `bson:"_id" json:"id"`
	Name string             `json:"name"`
}

type City struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `json:"name"`
	Shipment int                `json:"shipment"`
	Days     int                `json:"days"`
	Province *Province          `json:"province"`
}
