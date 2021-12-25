package controllers

import (
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yaien/clothes-store-api/pkg/api/helpers/response"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"github.com/yaien/clothes-store-api/pkg/api/services"
)

type CartController struct {
	Guests services.GuestService
	Items  services.ItemService
}

func (c *CartController) Add(w http.ResponseWriter, r *http.Request) {
	var data struct {
		ID       models.ID `json:"id"`
		Size     string    `json:"size"`
		Quantity int       `json:"quantity"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}

	product, err := c.Items.FindOneByID(r.Context(), data.ID)

	if err != nil {
		response.Error(w, errors.New("ITEM_NOT_FOUND"), http.StatusBadRequest)
		return
	}

	if product.Sizes == nil {
		response.Error(w, errors.New("ITEM_SOLD_OUT"), http.StatusBadRequest)
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

	guest := r.Context().Value("guest").(*models.Guest)

	item := &models.CartItem{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Quantity:    data.Quantity,
		Size:        size.Label,
	}

	if len(product.Pictures) > 0 {
		item.Picture = product.Pictures[0]
	}

	if guest.Cart == nil {
		guest.Cart = &models.Cart{}
	}
	if err := guest.Cart.AddItem(item); err != nil {
		response.Error(w, errors.New("PRODUCT_ALREADY_ADDED"), http.StatusBadRequest)
		return
	}
	if err = c.Guests.Update(r.Context(), guest); err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.Send(w, guest.Cart)
}

func (c *CartController) Remove(w http.ResponseWriter, r *http.Request) {
	guest := r.Context().Value("guest").(*models.Guest)
	itemID, err := primitive.ObjectIDFromHex(mux.Vars(r)["item_id"])
	if err != nil {
		response.Error(w, errors.New("INVALID_ITEM_ID"), http.StatusNotFound)
		return
	}

	if !guest.Cart.RemoveItem(itemID) {
		response.Error(w, errors.New("ITEM_NOT_FOUND"), http.StatusNotFound)
		return
	}
	if err := c.Guests.Update(r.Context(), guest); err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.Send(w, guest.Cart)
}
