package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yaien/clothes-store-api/api/models"
	"github.com/yaien/clothes-store-api/api/services"
	"github.com/yaien/clothes-store-api/core"
)

// GuestController -> guest controller
type GuestController struct {
	*core.Controller
	Guests services.GuestService
}

func (c *GuestController) Create(w http.ResponseWriter, r *http.Request) {
	guest := &models.Guest{}
	if err := c.Guests.Create(guest); err != nil {
		c.JSON(w, err, http.StatusInternalServerError)
		return
	}
	c.Send(w, guest)
}

func (c *GuestController) Param(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	id := mux.Vars(r)["guest_id"]
	guest, err := c.Guests.Get(id)
	if err != nil {
		c.Error(w, errors.New("GUEST_NOT_FOUND"), http.StatusNotFound)
		return
	}
	ctx := context.WithValue(r.Context(), core.Key("guest"), guest)
	log.Println(guest)
	next(w, r.WithContext(ctx))
}

func (c *GuestController) Show(w http.ResponseWriter, r *http.Request) {
	guest := r.Context().Value(core.Key("guest")).(*models.Guest)
	c.Send(w, guest)
}
