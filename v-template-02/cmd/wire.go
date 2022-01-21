//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"
	"net/http"
	"v-template-02/config"
	"v-template-02/internal/controller"
	"v-template-02/internal/repository"
	"v-template-02/internal/service"
	"v-template-02/router"
)

func InitApp(cfg *config.Config) (*http.Server, func(), error) {

	panic(wire.Build(repository.ProviderSet, service.ProviderSet, controller.ProviderSet, router.InitRouterSet, InitServerSet))
}
