package services

import (
	"context"
	"time"

	"github.com/yaien/clothes-store-api/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InvoiceService interface {
	Create(invoice *models.Invoice) error
}

type invoiceService struct {
	invoices *mongo.Collection
	index    IndexService
}

func (s *invoiceService) Create(invoice *models.Invoice) error {
	invoice.ID = primitive.NewObjectID()
	invoice.Ref = s.index.Get("invoice")
	invoice.CreatedAt = time.Now().Unix()
	invoice.Status = models.Pending
	_, err := s.invoices.InsertOne(context.TODO(), invoice)
	return err
}

func NewInvoiceService(db *mongo.Database) InvoiceService {
	return &invoiceService{
		invoices: db.Collection("invoices"),
		index:    NewIndexService(db),
	}
}
