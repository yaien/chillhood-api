package controller

import (
	"context"
	"errors"
	"github.com/yaien/clothes-store-api/pkg/entity"
	response "github.com/yaien/clothes-store-api/pkg/interface/rest/helpers/response"
	"github.com/yaien/clothes-store-api/pkg/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"

	"github.com/gorilla/mux"
)

// GuestController -> guest controller
type GuestController struct {
	Guests service.GuestService
}

func (c *GuestController) Create(w http.ResponseWriter, r *http.Request) {
	guest := &entity.Guest{}
	if err := c.Guests.Create(r.Context(), guest); err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.Send(w, guest)
}

func (c *GuestController) Param(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["guest_id"])
	if err != nil {
		response.Error(w, &entity.Error{Code: "INVALID_GUEST_ID", Err: err}, http.StatusBadRequest)
		return
	}
	guest, err := c.Guests.Get(r.Context(), id)
	if err != nil {
		response.Error(w, errors.New("GUEST_NOT_FOUND"), http.StatusNotFound)
		return
	}
	ctx := context.WithValue(r.Context(), "guest", guest)
	next(w, r.WithContext(ctx))
}

func (c *GuestController) Show(w http.ResponseWriter, r *http.Request) {
	guest := r.Context().Value("guest").(*entity.Guest)
	response.Send(w, guest)
}
