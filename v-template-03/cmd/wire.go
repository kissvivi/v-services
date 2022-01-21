//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"
	"net/http"
	"v-template-03/config"
	"v-template-03/internal/controller"
	"v-template-03/internal/repository"
	"v-template-03/internal/service"
	"v-template-03/router"
)

func InitApp(cfg *config.Config) (*http.Server, func(), error) {

	panic(wire.Build(repository.ProviderSet, service.ProviderSet, controller.ProviderSet, router.InitRouterSet, InitServerSet))
}
