package services

import (
	"context"
	"errors"
	"time"

	"github.com/gosimple/slug"
	"github.com/yaien/clothes-store-api/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemService interface {
	Create(product *models.Item) error
	Get(id string) (*models.Item, error)
	Update(product *models.Item) error
	Find() ([]*models.Item, error)
	Decrement(id string, size string, quantity int) error
	Increment(id string, size string, quantity int) error
}

type itemService struct {
	collection *mongo.Collection
}

func (p *itemService) Create(item *models.Item) error {
	item.ID = primitive.NewObjectID()
	item.CreatedAt = time.Now().Unix()
	item.Slug = slug.Make(item.Name)
	_, err := p.collection.InsertOne(context.TODO(), item)
	return err
}

func (p *itemService) Get(id string) (*models.Item, error) {
	var product models.Item
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = p.collection.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&product)
	return &product, err
}

func (p *itemService) Find() ([]*models.Item, error) {
	items := []*models.Item{}
	cursor, err := p.collection.Find(context.TODO(), bson.M{})
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

func (p *itemService) Update(product *models.Item) error {
	product.Slug = slug.Make(product.Name)
	filter := bson.M{"_id": product.ID}
	update := bson.M{"$set": product}
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
