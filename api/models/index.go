package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Index struct {
	ID    primitive.ObjectID `bson:"_id"`
	Key   string
	Value int
}
