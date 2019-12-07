package main

import "github.com/yaien/clothes-store-api/core"

import "log"

import "github.com/urfave/negroni"

import "github.com/yaien/clothes-store-api/api/routes"

func main() {
	app, err := core.NewApp()
	if err != nil {
		log.Fatal(err)
	}
	server := negroni.Classic()
	router := routes.Register(app)
	server.UseHandler(router)
	server.Run(app.Config.Address)
}
