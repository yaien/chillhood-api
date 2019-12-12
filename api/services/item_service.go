package services

import (
	"context"
	"time"

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
}

type itemService struct {
	collection *mongo.Collection
}

func (p *itemService) Create(product *models.Item) error {
	product.ID = primitive.NewObjectID()
	product.CreatedAt = time.Now().Unix()
	_, err := p.collection.InsertOne(context.TODO(), product)
	if err != nil {
		return err
	}
	return nil
}

func (p *itemService) Get(id string) (*models.Item, error) {
	var product models.Item
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = p.collection.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *itemService) Find() ([]*models.Item, error) {
	var items []*models.Item
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
	filter := bson.M{"_id": product.ID}
	update := bson.M{"$set": product}
	_, err := p.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func NewItemService(db *mongo.Database) ItemService {
	return &itemService{db.Collection("items")}
}
