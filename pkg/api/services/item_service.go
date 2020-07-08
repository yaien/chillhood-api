package services

import (
	"context"
	"errors"
	"time"

	"github.com/gosimple/slug"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemService interface {
	Create(product *models.Item) error
	Get(id string) (*models.Item, error)
	Update(product *models.Item) error
	Find(filter interface{}) ([]*models.Item, error)
	FindOne(filter interface{}) (*models.Item, error)
	Decrement(id string, size string, quantity int) error
	Increment(id string, size string, quantity int) error
}

type itemService struct {
	collection *mongo.Collection
}

func (p *itemService) Create(item *models.Item) error {
	item.ID = primitive.NewObjectID()
	item.CreatedAt = time.Now()
	item.Slug = slug.Make(item.Name)
	_, err := p.collection.InsertOne(context.TODO(), item)
	return err
}

func (p *itemService) Get(id string) (*models.Item, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": _id}
	return p.FindOne(filter)
}

func (p *itemService) FindOne(filter interface{}) (*models.Item, error) {
	var item models.Item
	err := p.collection.FindOne(context.TODO(), filter).Decode(&item)
	return &item, err
}

func (p *itemService) Find(filter interface{}) ([]*models.Item, error) {
	items := []*models.Item{}
	cursor, err := p.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var product models.Item
		err := cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		items = append(items, &product)
	}
	return items, nil
}

func (p *itemService) Update(item *models.Item) error {
	item.Slug = slug.Make(item.Name)
	item.UpdatedAt = time.Now()
	filter := bson.M{"_id": item.ID}
	update := bson.M{"$set": item}
	_, err := p.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (p *itemService) Decrement(id string, size string, quantity int) error {
	item, err := p.Get(id)
	if err != nil {
		return err
	}
	sz, err := item.Size(size)
	if err != nil {
		return err
	}
	if sz.Existence < quantity {
		return errors.New("INVALID_QUANTITY")
	}

	sz.Existence -= quantity
	return p.Update(item)
}

func (p *itemService) Increment(id string, size string, quantity int) error {
	item, err := p.Get(id)
	if err != nil {
		return err
	}
	sz, err := item.Size(size)
	if err != nil {
		return err
	}

	sz.Existence += quantity
	return p.Update(item)
}

func NewItemService(db *mongo.Database) ItemService {
	return &itemService{db.Collection("items")}
}
