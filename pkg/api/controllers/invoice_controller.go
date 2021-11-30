package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
	Cities   services.CityService
	Emails   services.EmailService
}

func (i *InvoiceController) Create(w http.ResponseWriter, r *http.Request) {
	var payload input.Invoice
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}

	guest := r.Context().Value("guest").(*models.Guest)

	cart, err := i.Carts.New(payload.Items)
	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}

	city, err := i.Cities.FindOne(r.Context(), models.FindOneCityOptions{
		Name:         payload.Shipping.City,
		ProvinceName: payload.Shipping.Province,
	})
	if err != nil {
		response.Error(w, fmt.Errorf("CITY_NOT_FOUND: %w", err), http.StatusBadRequest)
		return
	}

	cart.Shipping = city.Shipment
	cart.Refresh()

	invoice := &models.Invoice{Cart: cart, Shipping: payload.Shipping, GuestID: guest.ID}
	if err := i.Invoices.Create(r.Context(), invoice); err != nil {
		log.Println(err)
		response.Error(w, errors.New("SERVER_FAILED"), http.StatusInternalServerError)
		return
	}

	response.Send(w, invoice)
}

func (i *InvoiceController) Show(w http.ResponseWriter, r *http.Request) {
	invoice := r.Context().Value("invoice").(*models.Invoice)
	response.Send(w, invoice)
}

func (i *InvoiceController) Find(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	status := query.Get("status")
	search := query.Get("search")
	invoices, err := i.Invoices.Search(r.Context(), models.SearchInvoiceOptions{
		Query:  search,
		Status: models.InvoiceStatus(status),
	})
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	response.Send(w, invoices)
}

func (i *InvoiceController) GetByRef(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ref := mux.Vars(r)["invoice_ref"]
	invoice, err := i.Invoices.FindOneByRef(r.Context(), ref)
	if err != nil {
		response.Error(w, fmt.Errorf("INVOICE_NOT_FOUND: %s", err), http.StatusNotFound)
		return
	}
	ctx := context.WithValue(r.Context(), "invoice", invoice)
	next(w, r.WithContext(ctx))
}

func (i *InvoiceController) Get(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	id := mux.Vars(r)["invoice_id"]
	invoice, err := i.Invoices.FindOneByID(r.Context(), id)
	if err != nil {
		response.Error(w, fmt.Errorf("INVOICE_NOT_FOUND: %s", err), http.StatusNotFound)
		return
	}
	ctx := context.WithValue(r.Context(), "invoice", invoice)
	next(w, r.WithContext(ctx))
}

func (i *InvoiceController) SetTransport(w http.ResponseWriter, r *http.Request) {
	invoice := r.Context().Value("invoice").(*models.Invoice)
	if invoice.Status != models.Accepted {
		response.Error(w, errors.New("INVOICE_INVALID"), http.StatusBadRequest)
		return
	}

	var transport models.Transport
	err := json.NewDecoder(r.Body).Decode(&transport)
	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}
	invoice.Status = models.Completed
	invoice.Shipping.Status = models.Sended
	invoice.Shipping.Transport = &transport
	err = i.Invoices.Update(r.Context(), invoice)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	i.Emails.NotifyTransport(invoice)
	response.Send(w, invoice)
}
