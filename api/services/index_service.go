package services

import (
	"context"
	"strconv"

	"github.com/yaien/clothes-store-api/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IndexService interface {
	Get(key string) string
}

type indexService struct {
	indexes *mongo.Collection
}

func (s *indexService) Get(key string) string {
	var index models.Index
	filter := bson.M{"key": key}
	update := bson.M{
		"$inc": bson.M{
			"value": 1,
		},
	}
	s.indexes.UpdateOne(context.TODO(), filter, update)
	s.indexes.FindOne(context.TODO(), filter).Decode(&index)
	return strconv.Itoa(index.Value)
}

func NewIndexService(db *mongo.Database) IndexService {
	return &indexService{
		indexes: db.Collection("indexes"),
	}
}
