package mongodb

import (
	"context"
	"errors"
	"github.com/yaien/clothes-store-api/pkg/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProvinceRepository struct {
	collection *mongo.Collection
}

func (p *ProvinceRepository) Search(ctx context.Context, opts entity.SearchProvinceOptions) ([]*entity.Province, error) {
	var provinces []*entity.Province
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
		var province entity.Province
		err = cursor.Decode(&province)
		if err != nil {
			return nil, err
		}
		provinces = append(provinces, &province)
	}
	if provinces == nil {
		provinces = make([]*entity.Province, 0)
	}
	return provinces, nil
}

func (p *ProvinceRepository) FindOneByName(ctx context.Context, name string) (*entity.Province, error) {
	var prov entity.Province
	filter := bson.M{"name": name}
	err := p.collection.FindOne(ctx, filter).Decode(&prov)
	if err == nil {
		return &prov, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, &entity.Error{Code: "NOT_FOUND", Err: err}
	}
	return nil, err
}

func (p *ProvinceRepository) Create(ctx context.Context, pr *entity.Province) error {
	pr.ID = primitive.NewObjectID()
	_, err := p.collection.InsertOne(ctx, pr)
	return err
}

func NewProvinceRepository(db *mongo.Database) *ProvinceRepository {
	return &ProvinceRepository{collection: db.Collection("provinces")}
}
