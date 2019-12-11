package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InvoiceStatus string

type Invoice struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Ref       string             `json:"ref"`
	Cart      *Cart              `json:"cart"`
	Shipping  *Shipping          `json:"shipping"`
	Status    InvoiceStatus      `json:"status"`
	CreatedAt int64              `json:"createdAt"`
}

const (
	Accepted InvoiceStatus = "accepted"
	Rejected InvoiceStatus = "rejected"
	Pending  InvoiceStatus = "pending"
)
