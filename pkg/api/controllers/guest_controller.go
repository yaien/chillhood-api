package controllers

import (
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yaien/clothes-store-api/pkg/api/helpers/response"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"github.com/yaien/clothes-store-api/pkg/api/services"
)

// GuestController -> guest controller
type GuestController struct {
	Guests services.GuestService
}

func (c *GuestController) Create(w http.ResponseWriter, r *http.Request) {
	guest := &models.Guest{}
	if err := c.Guests.Create(r.Context(), guest); err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.Send(w, guest)
}

func (c *GuestController) Param(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	id := mux.Vars(r)["guest_id"]
	guest, err := c.Guests.Get(r.Context(), id)
	if err != nil {
		response.Error(w, errors.New("GUEST_NOT_FOUND"), http.StatusNotFound)
		return
	}
	ctx := context.WithValue(r.Context(), "guest", guest)
	next(w, r.WithContext(ctx))
}

func (c *GuestController) Show(w http.ResponseWriter, r *http.Request) {
	guest := r.Context().Value("guest").(*models.Guest)
	response.Send(w, guest)
}
