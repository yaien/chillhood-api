package service

import (
	"github.com/yaien/clothes-store-api/pkg/infrastructure"
	"github.com/yaien/clothes-store-api/pkg/interface/rest/helpers/epayco"
)

// ConfigService interface
type ConfigService interface {
	Cloudinary() *infrastructure.CloudinaryConfig
	Epayco() *epayco.CheckoutArgs
}

type configService struct {
	config *infrastructure.Config
}

func (cs *configService) Cloudinary() *infrastructure.CloudinaryConfig {
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
func NewConfigService(config *infrastructure.Config) ConfigService {
	return &configService{config}
}
