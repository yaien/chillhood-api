package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/yaien/clothes-store-api/api/controllers"
)

func invoice(router *mux.Router, mod *module) {
	invoice := &controllers.InvoiceController{
		Invoices: mod.service.invoices,
	}
	guest := &controllers.GuestController{
		Guests: mod.service.guests,
	}

	router.Handle("/api/v1/public/guests/{guest_id}/invoices", negroni.New(
		negroni.HandlerFunc(guest.Param),
		negroni.WrapFunc(invoice.Create),
	)).Methods("POST")

	router.HandleFunc("/api/v1/invoices/{invoice_ref}", invoice.Show).Methods("GET")
}
