package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/yaien/clothes-store-api/api/controllers"
	"github.com/yaien/clothes-store-api/api/services"
	"github.com/yaien/clothes-store-api/core"
)

func product(router *mux.Router, app *core.App) {
	c := &controllers.ProductController{
		Products: services.Product(app.DB),
	}

	router.HandleFunc("/api/v1/products", c.Create).Methods("POST")

	router.HandleFunc("/api/v1/products", c.Find).Methods("GET")

	router.Handle("/api/v1/products/{product_id}", negroni.New(
		negroni.HandlerFunc(c.Param),
		negroni.WrapFunc(c.Show),
	)).Methods("GET")

	router.Handle("/api/v1/products/{product_id}", negroni.New(
		negroni.HandlerFunc(c.Param),
		negroni.WrapFunc(c.Update),
	)).Methods("PUT")

}
