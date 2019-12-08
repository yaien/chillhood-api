package services

import (
	"context"

	"github.com/yaien/clothes-store-api/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService interface {
	Create(product *models.Product) error
	Get(id string) (*models.Product, error)
	Update(product *models.Product) error
	Find() ([]*models.Product, error)
}

type productService struct {
	collection *mongo.Collection
}

func (p *productService) Create(product *models.Product) error {
	product.ID = primitive.NewObjectID()
	_, err := p.collection.InsertOne(context.TODO(), product)
	if err != nil {
		return err
	}
	return nil
}

func (p *productService) Get(id string) (*models.Product, error) {
	var product models.Product
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

func (p *productService) Find() ([]*models.Product, error) {
	var products []*models.Product
	cursor, err := p.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (p *productService) Update(product *models.Product) error {
	filter := bson.M{"_id": product.ID}
	update := bson.M{"$set": product}
	_, err := p.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func Product(db *mongo.Database) ProductService {
	return &productService{db.Collection("products")}
}
