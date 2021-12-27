package input

import (
	"github.com/yaien/clothes-store-api/pkg/entity"
)

type Item struct {
	ID       entity.ID `json:"id"`
	Size     string    `json:"size"`
	Quantity int       `json:"quantity"`
}

type Invoice struct {
	Items    []*Item
	Shipping *entity.Shipping
}
