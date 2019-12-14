package services

import (
	"context"
	"github.com/teris-io/shortid"
	"strings"
	"time"

	"github.com/yaien/clothes-store-api/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InvoiceService interface {
	Create(invoice *models.Invoice) error
	Update(invoice *models.Invoice) error
	GetByRef(ref string) (*models.Invoice, error)
}

type invoiceService struct {
	invoices *mongo.Collection
}

func (s *invoiceService) Create(invoice *models.Invoice) error {
	invoice.ID = primitive.NewObjectID()
	invoice.Ref = strings.ToUpper(shortid.MustGenerate())
	invoice.CreatedAt = time.Now().Unix()
	invoice.Status = models.Created
	_, err := s.invoices.InsertOne(context.TODO(), invoice)
	return err
}

func (s *invoiceService) GetByRef(ref string) (*models.Invoice, error) {
	var invoice models.Invoice
	filter := bson.M{"ref": ref}
	err := s.invoices.FindOne(context.TODO(), filter).Decode(&invoice)
	return &invoice, err
}

func (s *invoiceService) Update(invoice *models.Invoice) error {
	filter := bson.M{"_id": invoice.ID}
	update := bson.M{"$set": invoice}
	_, err := s.invoices.UpdateOne(context.TODO(), filter, update)
	return err
}

func NewInvoiceService(db *mongo.Database) InvoiceService {
	return &invoiceService{db.Collection("invoices")}
}
