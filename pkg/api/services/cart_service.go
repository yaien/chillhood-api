package services

import (
	"fmt"

	"github.com/yaien/clothes-store-api/pkg/api/helpers/input"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartService interface {
	Execute(cart *models.Cart) error
	Revert(cart *models.Cart) error
	New(items []*input.Item) (*models.Cart, error)
}

type cartService struct {
	Items ItemService
}

func (c *cartService) Execute(cart *models.Cart) error {
	for _, cartItem := range cart.Items {
		err := c.Items.Decrement(cartItem.ID.Hex(), cartItem.Size, cartItem.Quantity)
		if err != nil {
			return err
		}
	}
	cart.Executed = true
	return nil
}

func (c *cartService) Revert(cart *models.Cart) error {
	for _, cartItem := range cart.Items {
		err := c.Items.Increment(cartItem.ID.Hex(), cartItem.Size, cartItem.Quantity)
		if err != nil {
			return err
		}
	}
	cart.Executed = false
	return nil
}

func (c *cartService) New(requests []*input.Item) (*models.Cart, error) {
	cart := &models.Cart{}
	for _, request := range requests {
		id, err := primitive.ObjectIDFromHex(request.ID)
		if err != nil {
			return nil, fmt.Errorf("INVALID_ITEM_ID: %s isn't a a valid object id: %w", request.ID, err)
		}
		item, err := c.Items.FindOne(bson.M{"_id": id, "active": true})
		if err != nil {
			return nil, fmt.Errorf("ITEM_NOT_FOUND: item %s doesn't exist or is inactive: %w", id, err)
		}
		size, err := item.Size(request.Size)
		if err != nil {
			return nil, fmt.Errorf("INVALID_SIZE: %w", err)
		}
		if request.Quantity > size.Existence {
			return nil, fmt.Errorf("SOLD_OUT: there is only %d %s (id: %s, size: %s) items, requested %d",
				size.Existence, item.Name, id, size.Label, request.Quantity)
		}

		var picture *models.Picture
		if len(item.Pictures) > 0 {
			picture = item.Pictures[0]
		}

		cart.Items = append(cart.Items, &models.CartItem{
			ID:       item.ID,
			Name:     item.Name,
			Price:    item.Price,
			Quantity: request.Quantity,
			Size:     request.Size,
			Picture:  picture,
		})
	}
	cart.Refresh()
	return cart, nil
}

func NewCartService(items ItemService) CartService {
	return &cartService{items}
}
