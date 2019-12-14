package services

import "github.com/yaien/clothes-store-api/api/models"

type CartService interface {
	Execute(cart *models.Cart) error
	Revert(cart *models.Cart) error
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

func NewCartService(items ItemService) CartService {
	return &cartService{items}
}
