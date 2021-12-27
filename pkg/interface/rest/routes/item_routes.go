package routes

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/yaien/clothes-store-api/pkg/interface/rest/controller"
)

func item(router *mux.Router, mod *module) {
	c := &controller.ItemController{
		Items: mod.service.items,
	}
	param := negroni.HandlerFunc(c.Param)

	router.HandleFunc("/api/v1/public/items", c.FindActive).Methods("GET")
	router.HandleFunc("/api/v1/public/items/{item_slug}", c.Slug).Methods("GET")

	router.Handle("/api/v1/items", negroni.New(
		mod.middleware.jwt,
		negroni.WrapFunc(c.Create),
	)).Methods("POST")

	router.Handle("/api/v1/items", negroni.New(
		mod.middleware.jwt,
		negroni.WrapFunc(c.Find),
	)).Methods("GET")

	router.Handle("/api/v1/items/{item_id}", negroni.New(
		mod.middleware.jwt,
		param,
		negroni.WrapFunc(c.Show),
	)).Methods("GET")

	router.Handle("/api/v1/items/{item_id}", negroni.New(
		mod.middleware.jwt,
		param,
		negroni.WrapFunc(c.Update),
	)).Methods("PUT")

}
