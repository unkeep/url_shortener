package service

import (
	"url_shortener/service/api"
	"url_shortener/service/database"
	"url_shortener/service/domain"
	"url_shortener/service/logger"

	"github.com/kelseyhightower/envconfig"
)

// Config is the service config
type Config struct {
	Port string `envconfig:"PORT" required:"true"`

	DB     database.Config
	Log    logger.Config
	Domain domain.Config
	API    api.Config
}

func (c *Config) Load() error {
	return envconfig.Process("", c)
}
