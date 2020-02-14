package input

import "github.com/yaien/clothes-store-api/api/models"

type Item struct {
	ID       string `json:"id"`
	Size     string `json:"size"`
	Quantity int    `json:"quantity"`
}

type Invoice struct {
	Items    []*Item
	Shipping *models.Shipping
}
