package main

import (
	"context"
	"log"
	"net/http"
	"net/url"

	"github.com/yaien/clothes-store-api/pkg/infrastructure/migrations"
	"github.com/yaien/clothes-store-api/pkg/interface/rest/routes"
	"github.com/yaien/ngrok"
	"github.com/yaien/p2p"

	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"github.com/urfave/negroni"
	"github.com/yaien/clothes-store-api/pkg/infrastructure"
)

func server() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start cloth store rest server",
		Run: func(cmd *cobra.Command, args []string) {

			app, err := infrastructure.NewApp()
			if err != nil {
				log.Fatal(err)
			}

			if app.Config.Ngrok.Enabled {
				tunnel, err := ngrok.Open(context.Background(), ngrok.Options{
					Addr:      app.Config.Address,
					AuthToken: app.Config.Ngrok.Token,
				})
				if err != nil {
					log.Fatalf("failed creating ngrok tunnel: %s", err)
				}

				defer tunnel.Close()

				app.Config.BaseURL, err = url.Parse(tunnel.Url())
				if err != nil {
					log.Fatalf("failed parsing ngrok url: %s", err)
				}

				app.P2P.SetCurrentAddr(tunnel.Url())

				log.Println("app base url was set to", app.Config.BaseURL)
			}

			go func() {
				updater := migrations.NewUpdater(app.DB)
				err := updater.Update(context.TODO())
				if err != nil {
					log.Fatalln(err)
				}
				app.P2P.Start()
			}()

			server := negroni.New(negroni.NewRecovery(), negroni.NewLogger())
			server.Use(cors.New(cors.Options{
				AllowedOrigins: app.Config.Client.Origins,
				AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE"},
			}))

			mx := http.NewServeMux()
			p2p.NewHttpServer(app.P2P, p2p.NewSubscriber(app.P2P.Channel()), app.Config.P2P.Key).Register(mx)
			server.UseFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
				h, p := mx.Handler(r)
				if p != "" {
					h.ServeHTTP(w, r)
					return
				}
				next(w, r)
			})

			router := routes.Register(app)
			server.UseHandler(router)
			server.Run(app.Config.Address)
		},
	}
}
