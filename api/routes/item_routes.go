package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/yaien/clothes-store-api/api/controllers"
	"github.com/yaien/clothes-store-api/api/services"
	"github.com/yaien/clothes-store-api/core"
)

func item(router *mux.Router, app *core.App) {
	c := &controllers.ItemController{
		Items: services.NewItemService(app.DB),
	}

	router.HandleFunc("/api/v1/items", c.Create).Methods("POST")

	router.HandleFunc("/api/v1/items", c.Find).Methods("GET")

	router.Handle("/api/v1/items/{item_id}", negroni.New(
		negroni.HandlerFunc(c.Param),
		negroni.WrapFunc(c.Show),
	)).Methods("GET")

	router.Handle("/api/v1/items/{item_id}", negroni.New(
		negroni.HandlerFunc(c.Param),
		negroni.WrapFunc(c.Update),
	)).Methods("PUT")

}
