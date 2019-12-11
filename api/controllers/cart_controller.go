package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yaien/clothes-store-api/api/helpers/response"
	"github.com/yaien/clothes-store-api/api/models"
	"github.com/yaien/clothes-store-api/api/services"
)

type payload struct {
	Product  string `json:"product"`
	Size     string `json:"size"`
	Quantity int    `json:"quantity"`
}

type Cart struct {
	Guests   services.GuestService
	Products services.ProductService
}

func (c *Cart) Add(w http.ResponseWriter, r *http.Request) {
	var data payload
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}

	product, err := c.Products.Get(data.Product)

	if err != nil {
		response.Error(w, errors.New("PRODUCT_NOT_FOUND"), http.StatusBadRequest)
		return
	}

	if product.Sizes == nil {
		response.Error(w, errors.New("PRODUCT_SOLD_OUT"), http.StatusBadRequest)
		return
	}

	size, err := product.Size(data.Size)

	if err != nil {
		response.Error(w, errors.New("SIZE_NOT_FOUND"), http.StatusBadRequest)
		return
	}

	if size.Existence < data.Quantity {
		response.Error(w, errors.New("SIZE_SOLD_OUT"), http.StatusBadRequest)
		return
	}

	guest := r.Context().Value(key("guest")).(*models.Guest)

	item := &models.Item{
		Product:  product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: data.Quantity,
		Size:     size.Label,
	}

	if guest.Cart == nil {
		guest.Cart = &models.Cart{}
	} else if guest.Cart.HasProduct(product.ID) {
		response.Error(w, errors.New("PRODUCT_ALREADY_ADDED"), http.StatusBadRequest)
		return
	}

	guest.Cart.Items = append(guest.Cart.Items, item)
	guest.Cart.Refresh()

	if err = c.Guests.Update(guest); err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	response.Send(w, guest.Cart)

}

func (c *Cart) Remove(w http.ResponseWriter, r *http.Request) {
	guest := r.Context().Value(key("guest")).(*models.Guest)
	product := mux.Vars(r)["product_id"]
	if err := guest.Cart.Remove(product); err != nil {
		response.Error(w, errors.New("PRODUCT_NOT_FOUND"), http.StatusNotFound)
		return
	}
	guest.Cart.Refresh()
	if err := c.Guests.Update(guest); err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.Send(w, guest.Cart)
}
