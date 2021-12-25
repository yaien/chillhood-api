package input

import "github.com/yaien/clothes-store-api/pkg/api/models"

type Item struct {
	ID       models.ID `json:"id"`
	Size     string    `json:"size"`
	Quantity int       `json:"quantity"`
}

type Invoice struct {
	Items    []*Item
	Shipping *models.Shipping
}
