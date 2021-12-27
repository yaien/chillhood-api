package entity

import (
	"context"
	epayco2 "github.com/yaien/clothes-store-api/pkg/interface/rest/helpers/epayco"
	"time"
)

type Invoice struct {
	ID        ID               `bson:"_id" json:"id"`
	Ref       string           `json:"ref"`
	Cart      *Cart            `json:"cart"`
	Shipping  *Shipping        `json:"shipping"`
	Status    InvoiceStatus    `json:"status"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
	Payment   *epayco2.Payment `json:"-"`
	GuestID   ID               `json:"guestId"`
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
	FindOneByID(ctx context.Context, id ID) (*Invoice, error)
	FindOneByRef(ctx context.Context, ref string) (*Invoice, error)
	Search(ctx context.Context, opts SearchInvoiceOptions) ([]*Invoice, error)
	Update(ctx context.Context, invoice *Invoice) error
}

type SearchInvoiceOptions struct {
	Query  string
	Status InvoiceStatus
}
