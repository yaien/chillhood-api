package controller

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/yaien/clothes-store-api/pkg/entity"
	response "github.com/yaien/clothes-store-api/pkg/interface/rest/helpers/response"
	"github.com/yaien/clothes-store-api/pkg/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"

	"github.com/gorilla/mux"
)

type ItemController struct {
	Items service.ItemService
}

func (p *ItemController) Create(w http.ResponseWriter, r *http.Request) {
	var item entity.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}

	err = p.Items.Create(r.Context(), &item)
	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}
	response.Send(w, item)
}

func (p *ItemController) Find(w http.ResponseWriter, r *http.Request) {
	items, err := p.Items.Find(r.Context())
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.Send(w, items)
}

func (p *ItemController) FindActive(w http.ResponseWriter, r *http.Request) {
	items, err := p.Items.FindActive(r.Context())
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.Send(w, items)
}

func (p *ItemController) Slug(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["item_slug"]
	item, err := p.Items.FindOneBySlug(r.Context(), slug)
	if err != nil {
		response.Error(w, errors.New("ITEM_NOT_FOUND"), http.StatusNotFound)
		return
	}
	response.Send(w, item)
}

func (p *ItemController) Param(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["item_id"])
	if err != nil {
		response.Error(w, &entity.Error{Code: "INVALID_ITEM_ID", Err: err}, http.StatusBadRequest)
		return
	}
	item, err := p.Items.FindOneByID(r.Context(), id)
	if err != nil {
		response.Error(w, errors.New("ITEM_NOT_FOUND"), http.StatusNotFound)
		return
	}
	ctx := context.WithValue(r.Context(), key("item"), item)
	next(w, r.WithContext(ctx))
}

func (p *ItemController) Show(w http.ResponseWriter, r *http.Request) {
	item := r.Context().Value(key("item")).(*entity.Item)
	response.Send(w, item)
}

func (p *ItemController) Update(w http.ResponseWriter, r *http.Request) {
	var data entity.Item
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}

	item := r.Context().Value(key("item")).(*entity.Item)
	data.ID = item.ID
	if err := p.Items.Update(r.Context(), &data); err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}
	response.Send(w, data)
}
