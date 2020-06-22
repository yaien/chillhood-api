package services

import (
	"github.com/yaien/clothes-store-api/pkg/api/helpers/epayco"
	"github.com/yaien/clothes-store-api/pkg/core"
)

// ConfigService interface
type ConfigService interface {
	Cloudinary() *core.ClodinaryConfig
	Epayco() *epayco.CheckoutArgs
}

type configService struct {
	config *core.Config
}

func (cs *configService) Cloudinary() *core.ClodinaryConfig {
	return cs.config.Cloudinary
}

func (cs *configService) Epayco() *epayco.CheckoutArgs {
	return &epayco.CheckoutArgs{
		Key:          cs.config.Epayco.PublicKey,
		Test:         cs.config.Epayco.Test,
		Response:     cs.config.BaseURL.String() + "/api/v1/public/epayco/response",
		Confirmation: cs.config.BaseURL.String() + "/api/v1/public/epayco/confirmation",
	}
}

// NewConfigService returns a new config service
func NewConfigService(config *core.Config) ConfigService {
	return &configService{config}
}
