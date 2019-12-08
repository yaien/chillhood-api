package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yaien/clothes-store-api/api/models"
	"github.com/yaien/clothes-store-api/api/services"
)

// Guest -> guest controller
type Guest struct {
	*controller
	Guests services.GuestService
}

func (c *Guest) Create(w http.ResponseWriter, r *http.Request) {
	guest := &models.Guest{}
	if err := c.Guests.Create(guest); err != nil {
		c.JSON(w, err, http.StatusInternalServerError)
		return
	}
	c.Send(w, guest)
}

func (c *Guest) Param(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	id := mux.Vars(r)["guest_id"]
	guest, err := c.Guests.Get(id)
	if err != nil {
		c.Error(w, errors.New("GUEST_NOT_FOUND"), http.StatusNotFound)
		return
	}
	ctx := context.WithValue(r.Context(), key("guest"), guest)
	log.Println(guest)
	next(w, r.WithContext(ctx))
}

func (c *Guest) Show(w http.ResponseWriter, r *http.Request) {
	guest := r.Context().Value(key("guest")).(*models.Guest)
	c.Send(w, guest)
}
