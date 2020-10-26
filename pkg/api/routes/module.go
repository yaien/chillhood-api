package routes

import (
	"github.com/urfave/negroni"
	"github.com/yaien/clothes-store-api/pkg/api/middlewares"
	"github.com/yaien/clothes-store-api/pkg/api/services"
	"github.com/yaien/clothes-store-api/pkg/core"
)

type service struct {
	users     services.UserService
	carts     services.CartService
	guests    services.GuestService
	items     services.ItemService
	epayco    services.EpaycoService
	invoices  services.InvoiceService
	tokens    services.TokenService
	config    services.ConfigService
	cities    services.CityService
	provinces services.ProvinceService
	slack     services.SlackService
	email     services.EmailService
}

type middleware struct {
	jwt negroni.Handler
}

type module struct {
	service    *service
	middleware *middleware
}

func bundle(app *core.App) *module {
	email := services.NewEmailService(app.Config.SMTP, app.Templates)
	cities := services.NewCityService(app.DB)
	provinces := services.NewProvinceService(app.DB)
	items := services.NewItemService(app.DB)
	users := services.NewUserService(app.DB)
	carts := services.NewCartService(items)
	guests := services.NewGuestService(app.DB)
	invoices := services.NewInvoiceService(app.DB)
	slack := services.NewSlackService(app.Slack, app.Config.Slack)
	epayco := services.NewEpaycoService(app.Config.Epayco, app.Config.BaseURL, invoices, carts, guests, slack, email)
	tokens := services.NewTokenService(app.Config.Client, app.Config.JWT, users)
	config := services.NewConfigService(app.Config)

	return &module{
		service: &service{
			cities:    cities,
			provinces: provinces,
			users:     users,
			carts:     carts,
			guests:    guests,
			items:     items,
			epayco:    epayco,
			invoices:  invoices,
			tokens:    tokens,
			config:    config,
			slack:     slack,
			email:     email,
		},
		middleware: &middleware{
			jwt: &middlewares.JWTGuard{Tokens: tokens, Users: users},
		},
	}
}
