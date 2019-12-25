package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/yaien/clothes-store-api/api/controllers"
)

func item(router *mux.Router, mod *module) {
	c := &controllers.ItemController{
		Items: mod.service.items,
	}

	router.HandleFunc("/api/v1/public/items", c.FindActive).Methods("GET")
	router.HandleFunc("/api/v1/public/items/{item_slug}", c.Slug).Methods("GET")

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
