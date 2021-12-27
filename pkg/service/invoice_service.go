package service

import (
	"context"
	"github.com/yaien/clothes-store-api/pkg/entity"
	"strings"
	"time"

	"github.com/teris-io/shortid"
)

type InvoiceService interface {
	Create(ctx context.Context, invoice *entity.Invoice) error
	Update(ctx context.Context, invoice *entity.Invoice) error
	FindOneByID(ctx context.Context, id entity.ID) (*entity.Invoice, error)
	FindOneByRef(ctx context.Context, ref string) (*entity.Invoice, error)
	Search(ctx context.Context, opts entity.SearchInvoiceOptions) ([]*entity.Invoice, error)
}

type invoiceService struct {
	invoices entity.InvoiceRepository
}

func (s *invoiceService) Create(ctx context.Context, invoice *entity.Invoice) error {
	invoice.Ref = strings.ToUpper(shortid.MustGenerate())
	invoice.CreatedAt = time.Now()
	invoice.UpdatedAt = time.Now()
	invoice.Status = entity.Created
	return s.invoices.Create(ctx, invoice)
}

func (s *invoiceService) FindOneByID(ctx context.Context, id entity.ID) (*entity.Invoice, error) {
	return s.invoices.FindOneByID(ctx, id)
}

func (s *invoiceService) FindOneByRef(ctx context.Context, ref string) (*entity.Invoice, error) {
	return s.invoices.FindOneByRef(ctx, ref)
}

func (s *invoiceService) Search(ctx context.Context, opts entity.SearchInvoiceOptions) ([]*entity.Invoice, error) {
	return s.invoices.Search(ctx, opts)
}

func (s *invoiceService) Update(ctx context.Context, invoice *entity.Invoice) error {
	invoice.UpdatedAt = time.Now()
	return s.invoices.Update(ctx, invoice)
}

func NewInvoiceService(invoices entity.InvoiceRepository) InvoiceService {
	return &invoiceService{invoices}
}
