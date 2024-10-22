package core

import (
	"github.com/aSel1x/Gin_Template/adapters"
	"github.com/aSel1x/Gin_Template/usecases"
)

type AppProvider struct {
	*Config
	*adapters.Adapters
	*usecases.Usecases
}

func NewAppProvider(config *Config) *AppProvider {
	newAdapters := adapters.NewAdapters(config.Postgres.DSN())
	newUsecases := usecases.NewUsecases(*newAdapters, config.App.Secret)

	return &AppProvider{
		Config:   config,
		Adapters: newAdapters,
		Usecases: newUsecases,
	}
}
