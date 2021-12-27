package routes

import (
	"github.com/urfave/negroni"
	"github.com/yaien/clothes-store-api/pkg/infrastructure"
	"github.com/yaien/clothes-store-api/pkg/interface/mongodb"
	middlewares2 "github.com/yaien/clothes-store-api/pkg/interface/rest/middlewares"
	"github.com/yaien/clothes-store-api/pkg/service"
)

type services struct {
	users     service.UserService
	carts     service.CartService
	guests    service.GuestService
	items     service.ItemService
	epayco    service.EpaycoService
	invoices  service.InvoiceService
	tokens    service.TokenService
	config    service.ConfigService
	cities    service.CityService
	provinces service.ProvinceService
	slack     service.SlackService
	email     service.EmailService
}

type middleware struct {
	jwt negroni.Handler
}

type module struct {
	service    *services
	middleware *middleware
}

func bundle(app *infrastructure.App) *module {
	email := service.NewEmailService(app.Config.SMTP, app.Templates)
	cities := service.NewCityService(mongodb.NewCityRepository(app.DB))
	provinces := service.NewProvinceService(mongodb.NewProvinceRepository(app.DB))
	items := service.NewItemService(mongodb.NewItemRepository(app.DB))
	users := service.NewUserService(mongodb.NewUserRepository(app.DB))
	carts := service.NewCartService(items)
	guests := service.NewGuestService(mongodb.NewGuestRepository(app.DB))
	invoices := service.NewInvoiceService(mongodb.NewMongoInvoiceRepository(app.DB))
	slack := service.NewSlackService(app.Slack, app.Config.Slack)
	epayco := service.NewEpaycoService(app.Config.Epayco, app.Config.BaseURL, invoices, carts, guests, slack, email)
	tokens := service.NewTokenService(app.Config.Client, app.Config.JWT, users)
	config := service.NewConfigService(app.Config)

	return &module{
		service: &services{
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
			jwt: &middlewares2.JWTGuard{Tokens: tokens, Users: users},
		},
	}
}
