package service

import (
	"context"
	"fmt"
	"github.com/yaien/clothes-store-api/pkg/entity"
	"github.com/yaien/clothes-store-api/pkg/interface/rest/helpers/input"
)

type CartService interface {
	Execute(cart *entity.Cart) error
	Revert(cart *entity.Cart) error
	New(items []*input.Item) (*entity.Cart, error)
}

type cartService struct {
	Items ItemService
}

func (c *cartService) Execute(cart *entity.Cart) error {
	for _, cartItem := range cart.Items {
		err := c.Items.Decrement(context.TODO(), cartItem.ID, cartItem.Size, cartItem.Quantity)
		if err != nil {
			return err
		}
	}
	cart.Executed = true
	return nil
}

func (c *cartService) Revert(cart *entity.Cart) error {
	for _, cartItem := range cart.Items {
		err := c.Items.Increment(context.TODO(), cartItem.ID, cartItem.Size, cartItem.Quantity)
		if err != nil {
			return err
		}
	}
	cart.Executed = false
	return nil
}

func (c *cartService) New(requests []*input.Item) (*entity.Cart, error) {
	cart := &entity.Cart{}
	for _, request := range requests {
		item, err := c.Items.FindOneActiveByID(context.TODO(), request.ID)
		if err != nil {
			return nil, fmt.Errorf("ITEM_NOT_FOUND: item %s doesn't exist or is inactive: %w", request.ID, err)
		}
		size, err := item.Size(request.Size)
		if err != nil {
			return nil, fmt.Errorf("INVALID_SIZE: %w", err)
		}
		if request.Quantity > size.Existence {
			return nil, fmt.Errorf("SOLD_OUT: there is only %d %s (id: %s, size: %s) items, requested %d",
				size.Existence, item.Name, request.ID, size.Label, request.Quantity)
		}

		var picture *entity.Picture
		if len(item.Pictures) > 0 {
			picture = item.Pictures[0]
		}

		cart.Items = append(cart.Items, &entity.CartItem{
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
