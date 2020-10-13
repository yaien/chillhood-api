package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/yaien/clothes-store-api/pkg/api/controllers"
)

func invoice(router *mux.Router, mod *module) {
	invoice := &controllers.InvoiceController{
		Invoices: mod.service.invoices,
		Carts:    mod.service.carts,
		Cities:   mod.service.cities,
	}
	guest := &controllers.GuestController{
		Guests: mod.service.guests,
	}

	router.Handle("/api/v1/public/guests/{guest_id}/invoices", negroni.New(
		negroni.HandlerFunc(guest.Param),
		negroni.WrapFunc(invoice.Create),
	)).Methods("POST")

	router.Handle("/api/v1/public/invoices/{invoice_ref}", negroni.New(
		negroni.HandlerFunc(invoice.GetByRef),
		negroni.WrapFunc(invoice.Show),
	)).Methods("GET")

	router.Handle("/api/v1/invoices", negroni.New(
		mod.middleware.jwt,
		negroni.WrapFunc(invoice.Find),
	)).Methods("GET")

	router.Handle("/api/v1/invoices/{invoice_id}", negroni.New(
		mod.middleware.jwt,
		negroni.HandlerFunc(invoice.Get),
		negroni.WrapFunc(invoice.Show),
	)).Methods("GET")

	router.Handle("/api/v1/invoices/{invoice_id}/transport", negroni.New(
		mod.middleware.jwt,
		negroni.HandlerFunc(invoice.Get),
		negroni.WrapFunc(invoice.SetTransport),
	)).Methods("PATCH")
}
