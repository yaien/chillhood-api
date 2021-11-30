package services

import (
	"context"
	"strings"
	"time"

	"github.com/teris-io/shortid"

	"github.com/yaien/clothes-store-api/pkg/api/models"
)

type InvoiceService interface {
	Create(ctx context.Context, invoice *models.Invoice) error
	Update(ctx context.Context, invoice *models.Invoice) error
	FindOneByID(ctx context.Context, id string) (*models.Invoice, error)
	FindOneByRef(ctx context.Context, ref string) (*models.Invoice, error)
	Search(ctx context.Context, opts models.SearchInvoiceOptions) ([]*models.Invoice, error)
}

type invoiceService struct {
	invoices models.InvoiceRepository
}

func (s *invoiceService) Create(ctx context.Context, invoice *models.Invoice) error {
	invoice.Ref = strings.ToUpper(shortid.MustGenerate())
	invoice.CreatedAt = time.Now()
	invoice.UpdatedAt = time.Now()
	invoice.Status = models.Created
	return s.invoices.Create(ctx, invoice)
}

func (s *invoiceService) FindOneByID(ctx context.Context, id string) (*models.Invoice, error) {
	return s.invoices.FindOneByID(ctx, id)
}

func (s *invoiceService) FindOneByRef(ctx context.Context, ref string) (*models.Invoice, error) {
	return s.invoices.FindOneByRef(ctx, ref)
}

func (s *invoiceService) Search(ctx context.Context, opts models.SearchInvoiceOptions) ([]*models.Invoice, error) {
	return s.invoices.Search(ctx, opts)
}

func (s *invoiceService) Update(ctx context.Context, invoice *models.Invoice) error {
	invoice.UpdatedAt = time.Now()
	return s.invoices.Update(ctx, invoice)
}

func NewInvoiceService(invoices models.InvoiceRepository) InvoiceService {
	return &invoiceService{invoices}
}
