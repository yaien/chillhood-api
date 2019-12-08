package controllers

import "github.com/yaien/clothes-store-api/api/services"

import "net/http"

import "encoding/json"

import "errors"

import "github.com/yaien/clothes-store-api/api/models"

type payload struct {
	product  string
	size     string
	quantity int
}

type Cart struct {
	*controller
	Guests   services.GuestService
	Products services.ProductService
}

func (c *Cart) AddProduct(w http.ResponseWriter, r *http.Request) {
	var data payload
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		c.Error(w, err, http.StatusBadRequest)
		return
	}

	product, err := c.Products.Get(data.product)

	if err != nil {
		c.Error(w, errors.New("PRODUCT_NOT_FOUND"), http.StatusBadRequest)
		return
	}

	if product.Sizes == nil {
		c.Error(w, errors.New("PRODUCT_SOLD_OUT"), http.StatusBadRequest)
		return
	}

	size, err := product.Size(data.size)

	if err != nil {
		c.Error(w, errors.New("SIZE_NOT_FOUND"), http.StatusBadRequest)
		return
	}

	if size.Existence < data.quantity {
		c.Error(w, errors.New("SIZE_SOLD_OUT"), http.StatusBadRequest)
		return
	}

	guest := r.Context().Value(key("guest")).(*models.Guest)

	item := &models.Item{
		Product:  product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: data.quantity,
		Size:     size.Label,
	}

	if guest.Cart == nil {
		guest.Cart = &models.Cart{}
	}

	guest.Cart.Items = append(guest.Cart.Items, item)
	guest.Cart.Compute()

	if err = c.Guests.Update(guest); err != nil {
		c.Error(w, err, http.StatusInternalServerError)
	}

	c.Send(w, guest.Cart)

}
