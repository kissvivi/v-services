//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"
	"net/http"
	"v-template-etcd/config"
	"v-template-etcd/internal/controller"
	"v-template-etcd/internal/repository"
	"v-template-etcd/internal/service"
	"v-template-etcd/router"
)

func InitApp(cfg *config.Config) (*http.Server, func(), error) {

	panic(wire.Build(repository.ProviderSet, service.ProviderSet, controller.ProviderSet, router.InitRouterSet, InitServerSet))
}
