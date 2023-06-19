package config

import (
	"strings"

	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/config"
)

type Config struct {
	AppOptions           AppOptions           `mapstructure:"appOptions"`
	MongoDocumentOptions MongoDocumentOptions `mapstructure:"mongoDocumentOptions"`
}

func NewConfig(env config.Environment) (*Config, error) {
	cfg, err := config.BindConfig[*Config](env)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

type AppOptions struct {
	DeliveryType string `mapstructure:"deliveryType"`
	ServiceName  string `mapstructure:"serviceName"`
}

func (cfg *AppOptions) GetMicroserviceNameUpper() string {
	return strings.ToUpper(cfg.ServiceName)
}

func (cfg *AppOptions) GetMicroserviceName() string {
	return cfg.ServiceName
}

type MongoDocumentOptions struct {
	Orders       string `mapstructure:"orders"       validate:"required" env:"Orders"`
	DatabaseName string `mapstructure:"databaseName" validate:"required" env:"DatabaseName"`
}
