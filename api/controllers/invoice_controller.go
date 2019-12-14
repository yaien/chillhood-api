package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yaien/clothes-store-api/api/helpers/response"
	"github.com/yaien/clothes-store-api/api/models"
	"github.com/yaien/clothes-store-api/api/services"
)

type InvoiceController struct {
	Invoices services.InvoiceService
}

func (i *InvoiceController) Create(w http.ResponseWriter, r *http.Request) {
	var shipping models.Shipping
	if err := json.NewDecoder(r.Body).Decode(&shipping); err != nil {
		response.Error(w, err, http.StatusBadGateway)
		return
	}
	guest := r.Context().Value(key("guest")).(*models.Guest)
	if guest.Cart == nil || len(guest.Cart.Items) == 0 {
		response.Error(w, errors.New("INVALID_CART"), http.StatusBadRequest)
		return
	}

	invoice := &models.Invoice{Cart: guest.Cart, Shipping: &shipping}
	if err := i.Invoices.Create(invoice); err != nil {
		log.Println(err)
		response.Error(w, errors.New("SERVER_FAILED"), http.StatusInternalServerError)
		return
	}

	response.Send(w, invoice)
}

func (i *InvoiceController) Show(w http.ResponseWriter, r *http.Request) {
	ref := mux.Vars(r)["invoice_ref"]
	invoice, err := i.Invoices.GetByRef(ref)
	if err != nil {
		response.Error(w, errors.New("INVOICE_NOT_FOUND"), http.StatusNotFound)
		return
	}
	response.Send(w, invoice)
}
