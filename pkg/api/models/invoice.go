package models

import (
	"context"
	"time"

	"github.com/yaien/clothes-store-api/pkg/api/helpers/epayco"
)

type Invoice struct {
	ID        string          `bson:"_id" json:"id"`
	Ref       string          `json:"ref"`
	Cart      *Cart           `json:"cart"`
	Shipping  *Shipping       `json:"shipping"`
	Status    InvoiceStatus   `json:"status"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	Payment   *epayco.Payment `json:"-"`
	GuestID   string          `json:"guestId"`
}

type InvoiceStatus string

const (
	Created   InvoiceStatus = "created"
	Accepted  InvoiceStatus = "accepted"
	Rejected  InvoiceStatus = "rejected"
	Pending   InvoiceStatus = "pending"
	Completed InvoiceStatus = "completed"
)

type InvoiceRepository interface {
	Create(ctx context.Context, invoice *Invoice) error
	FindOneByID(ctx context.Context, id string) (*Invoice, error)
	FindOneByRef(ctx context.Context, ref string) (*Invoice, error)
	Search(ctx context.Context, opts InvoiceSearchOptions) ([]*Invoice, error)
	Update(ctx context.Context, invoice *Invoice) error
}

type InvoiceSearchOptions struct {
	Query  string
	Status InvoiceStatus
}
