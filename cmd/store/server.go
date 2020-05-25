package main

import (
	"log"

	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"github.com/urfave/negroni"
	"github.com/yaien/clothes-store-api/pkg/api/routes"
	"github.com/yaien/clothes-store-api/pkg/core"
)

func server() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start cloth store api server",
		Run: func(cmd *cobra.Command, args []string) {
			app, err := core.NewApp()
			if err != nil {
				log.Fatal(err)
			}
			server := negroni.Classic()
			server.Use(cors.New(cors.Options{
				AllowedOrigins: app.Config.Client.Origins,
				AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE"},
			}))
			router := routes.Register(app)
			server.UseHandler(router)
			server.Run(app.Config.Address)
		},
	}
}
