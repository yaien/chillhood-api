package mongodb

import (
	"context"
	"errors"
	"github.com/yaien/clothes-store-api/pkg/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemRepository struct {
	collection *mongo.Collection
}

func (m *ItemRepository) Create(ctx context.Context, item *entity.Item) error {
	item.ID = primitive.NewObjectID()
	_, err := m.collection.InsertOne(ctx, item)
	return err
}

func (m *ItemRepository) CountByName(ctx context.Context, name string) (int64, error) {
	return m.collection.CountDocuments(ctx, bson.M{"name": name})
}

func (m *ItemRepository) CountByNameIgnore(ctx context.Context, id primitive.ObjectID, name string) (int64, error) {
	return m.collection.CountDocuments(ctx, bson.M{"name": name, "id": bson.M{"$ne": id}})
}

func (m *ItemRepository) FindOneByID(ctx context.Context, id primitive.ObjectID) (*entity.Item, error) {
	var item entity.Item
	err := m.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&item)
	if err == nil {
		return &item, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, &entity.Error{Code: "NOT_FOUND", Err: err}
	}
	return nil, err
}

func (m *ItemRepository) FindOneActiveByID(ctx context.Context, id primitive.ObjectID) (*entity.Item, error) {
	var item entity.Item
	err := m.collection.FindOne(ctx, bson.M{"_id": id, "active": true}).Decode(&item)
	if err == nil {
		return &item, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, &entity.Error{Code: "NOT_FOUND", Err: err}
	}
	return nil, err
}

func (m *ItemRepository) FindOneBySlug(ctx context.Context, slug string) (*entity.Item, error) {
	var item entity.Item
	err := m.collection.FindOne(ctx, bson.M{"slug": slug}).Decode(&item)
	if err == nil {
		return &item, nil
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, &entity.Error{Code: "NOT_FOUND", Err: err}
	}
	return nil, err
}

func (m *ItemRepository) Find(ctx context.Context) ([]*entity.Item, error) {
	var items []*entity.Item
	cursor, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var item entity.Item
		err := cursor.Decode(&item)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	if items == nil {
		items = []*entity.Item{}
	}
	return items, nil
}

func (m *ItemRepository) FindActive(ctx context.Context) ([]*entity.Item, error) {
	var items []*entity.Item
	cursor, err := m.collection.Find(ctx, bson.M{"active": true})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var item entity.Item
		err := cursor.Decode(&item)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	if items == nil {
		items = []*entity.Item{}
	}
	return items, nil
}

func (m *ItemRepository) Update(ctx context.Context, item *entity.Item) error {
	_, err := m.collection.UpdateOne(ctx, bson.M{"_id": item.ID}, bson.M{"$set": item})
	return err
}

func NewItemRepository(db *mongo.Database) *ItemRepository {
	return &ItemRepository{db.Collection("items")}
}
