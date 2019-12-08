package routes

import "github.com/gorilla/mux"

import "github.com/yaien/clothes-store-api/core"

import "github.com/yaien/clothes-store-api/api/controllers"

import "github.com/yaien/clothes-store-api/api/services"

import "github.com/urfave/negroni"

func product(router *mux.Router, app *core.App) {
	c := &controllers.Product{
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
