package repository

import (
	"context"
	"errors"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProvinceRepository struct {
	collection *mongo.Collection
}

func (p *MongoProvinceRepository) FindOneByName(ctx context.Context, name string) (*models.Province, error) {
	var prov models.Province
	filter := bson.M{"name": name}
	err := p.collection.FindOne(ctx, filter).Decode(&prov)
	if err == nil {
		return &prov, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, &models.Error{Code: "not_found", Err: err}
	}
	return nil, err
}

func (p *MongoProvinceRepository) Create(ctx context.Context, pr *models.Province) error {
	pr.ID = primitive.NewObjectID().Hex()
	_, err := p.collection.InsertOne(ctx, pr)
	return err
}

func NewMongoProvinceRepository(db *mongo.Database) *MongoProvinceRepository {
	return &MongoProvinceRepository{collection: db.Collection("provinces")}
}
