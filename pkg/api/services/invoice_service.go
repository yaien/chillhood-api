package services

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"

	"github.com/teris-io/shortid"

	"github.com/yaien/clothes-store-api/pkg/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InvoiceService interface {
	Create(invoice *models.Invoice) error
	Update(invoice *models.Invoice) error
	Get(id string) (*models.Invoice, error)
	Find(filter map[string]interface{}) ([]*models.Invoice, error)
	GetByRef(ref string) (*models.Invoice, error)
}

type invoiceService struct {
	invoices *mongo.Collection
}

func (s *invoiceService) Create(invoice *models.Invoice) error {
	invoice.ID = primitive.NewObjectID()
	invoice.Ref = strings.ToUpper(shortid.MustGenerate())
	invoice.CreatedAt = time.Now()
	invoice.UpdatedAt = time.Now()
	invoice.Status = models.Created
	_, err := s.invoices.InsertOne(context.TODO(), invoice)
	return err
}

func (s *invoiceService) Get(id string) (*models.Invoice, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var invoice models.Invoice
	filter := bson.M{"_id": _id}
	err = s.invoices.FindOne(context.TODO(), filter).Decode(&invoice)
	return &invoice, err
}

func (s *invoiceService) GetByRef(ref string) (*models.Invoice, error) {
	var invoice models.Invoice
	filter := bson.M{"ref": ref}
	err := s.invoices.FindOne(context.TODO(), filter).Decode(&invoice)
	return &invoice, err
}

func (s *invoiceService) Find(filter map[string]interface{}) ([]*models.Invoice, error) {
	var invoices []*models.Invoice
	cursor, err := s.invoices.Find(context.TODO(), filter, &options.FindOptions{
		Sort: bson.D{{"createdat", -1}},
	})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var invoice models.Invoice
		err := cursor.Decode(&invoice)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, &invoice)
	}
	if invoices == nil {
		invoices = make([]*models.Invoice, 0)
	}
	return invoices, nil
}

func (s *invoiceService) Update(invoice *models.Invoice) error {
	invoice.UpdatedAt = time.Now()
	filter := bson.M{"_id": invoice.ID}
	update := bson.M{"$set": invoice}
	_, err := s.invoices.UpdateOne(context.TODO(), filter, update)
	return err
}

func NewInvoiceService(db *mongo.Database) InvoiceService {
	return &invoiceService{db.Collection("invoices")}
}
