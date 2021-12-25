package repository

import (
	"context"
	"errors"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoItemRepository struct {
	collection *mongo.Collection
}

func (m *MongoItemRepository) Create(ctx context.Context, item *models.Item) error {
	item.ID = primitive.NewObjectID()
	_, err := m.collection.InsertOne(ctx, item)
	return err
}

func (m *MongoItemRepository) CountByName(ctx context.Context, name string) (int64, error) {
	return m.collection.CountDocuments(ctx, bson.M{"name": name})
}

func (m *MongoItemRepository) FindOneByID(ctx context.Context, id primitive.ObjectID) (*models.Item, error) {
	var item models.Item
	err := m.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&item)
	if err == nil {
		return &item, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, &models.Error{Code: "NOT_FOUND", Err: err}
	}
	return nil, err
}

func (m *MongoItemRepository) FindOneActiveByID(ctx context.Context, id primitive.ObjectID) (*models.Item, error) {
	var item models.Item
	err := m.collection.FindOne(ctx, bson.M{"_id": id, "active": true}).Decode(&item)
	if err == nil {
		return &item, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, &models.Error{Code: "NOT_FOUND", Err: err}
	}
	return nil, err
}

func (m *MongoItemRepository) FindOneBySlug(ctx context.Context, slug string) (*models.Item, error) {
	var item models.Item
	err := m.collection.FindOne(ctx, bson.M{"slug": slug}).Decode(&item)
	if err == nil {
		return &item, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, &models.Error{Code: "NOT_FOUND", Err: err}
	}
	return nil, err
}

func (m *MongoItemRepository) Find(ctx context.Context) ([]*models.Item, error) {
	var items []*models.Item
	cursor, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var item models.Item
		err := cursor.Decode(&item)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	return items, nil
}

func (m *MongoItemRepository) FindActive(ctx context.Context) ([]*models.Item, error) {
	var items []*models.Item
	cursor, err := m.collection.Find(ctx, bson.M{"active": true})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var item models.Item
		err := cursor.Decode(&item)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	return items, nil
}

func (m *MongoItemRepository) Update(ctx context.Context, item *models.Item) error {
	_, err := m.collection.UpdateOne(ctx, bson.M{"_id": item.ID}, bson.M{"$set": item})
	return err
}

func NewMongoItemRepository(db *mongo.Database) *MongoItemRepository {
	return &MongoItemRepository{db.Collection("items")}
}
