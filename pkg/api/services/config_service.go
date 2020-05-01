package services

import "github.com/yaien/clothes-store-api/pkg/core"

// ConfigService interface
type ConfigService interface {
	Cloudinary() *core.ClodinaryConfig
}

type configService struct {
	config *core.Config
}

func (cs *configService) Cloudinary() *core.ClodinaryConfig {
	return cs.config.Cloudinary
}

// NewConfigService returns a new config service
func NewConfigService(config *core.Config) ConfigService {
	return &configService{}
}
