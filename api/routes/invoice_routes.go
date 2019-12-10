package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/yaien/clothes-store-api/api/controllers"
	"github.com/yaien/clothes-store-api/api/services"
	"github.com/yaien/clothes-store-api/core"
)

func invoice(router *mux.Router, app *core.App) {
	invoice := &controllers.Invoice{
		Invoices: services.NewInvoiceService(app.DB),
	}
	guest := &controllers.Guest{
		Guests: services.Guest(app.DB),
	}

	router.Handle("/api/v1/guests/{guest_id}/invoice", negroni.New(
		negroni.HandlerFunc(guest.Param),
		negroni.WrapFunc(invoice.Create),
	)).Methods("POST")

}
