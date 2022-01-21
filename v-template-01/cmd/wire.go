//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"
	"net/http"
	"v-template-01/config"
	"v-template-01/internal/controller"
	"v-template-01/internal/repository"
	"v-template-01/internal/service"
	"v-template-01/router"
)

func InitApp(cfg *config.Config) (*http.Server, func(), error) {

	panic(wire.Build(repository.ProviderSet, service.ProviderSet, controller.ProviderSet, router.InitRouterSet, InitServerSet))
}
