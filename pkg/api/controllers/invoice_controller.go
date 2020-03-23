package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yaien/clothes-store-api/pkg/api/helpers/input"
	"github.com/yaien/clothes-store-api/pkg/api/helpers/response"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"github.com/yaien/clothes-store-api/pkg/api/services"
)

type InvoiceController struct {
	Invoices services.InvoiceService
	Carts    services.CartService
}

func (i *InvoiceController) Create(w http.ResponseWriter, r *http.Request) {
	var payload input.Invoice
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.Error(w, err, http.StatusBadGateway)
		return
	}

	cart, err := i.Carts.New(payload.Items)

	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}

	invoice := &models.Invoice{Cart: cart, Shipping: payload.Shipping}
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
