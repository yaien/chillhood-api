package repository

import (
	"context"
	"errors"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoProvinceRepository struct {
	collection *mongo.Collection
}

func (p *MongoProvinceRepository) Search(ctx context.Context, opts models.SearchProvinceOptions) ([]*models.Province, error) {
	var provinces []*models.Province
	filter := bson.M{
		"name": primitive.Regex{Pattern: opts.Name, Options: "i"},
	}
	skip := opts.Skip
	if skip < 0 {
		skip = 0
	}
	limit := opts.Limit
	if limit < 1 || limit > 100 {
		limit = 10
	}
	ops := options.Find().SetSkip(skip).SetLimit(limit)
	cursor, err := p.collection.Find(ctx, filter, ops)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var province models.Province
		err = cursor.Decode(&province)
		if err != nil {
			return nil, err
		}
		provinces = append(provinces, &province)
	}
	if provinces == nil {
		provinces = make([]*models.Province, 0)
	}
	return provinces, nil
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
