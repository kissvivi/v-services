// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"net/http"
	"v-template-etcd/config"
	"v-template-etcd/internal/controller"
	"v-template-etcd/internal/repository"
	"v-template-etcd/internal/service"
	"v-template-etcd/router"
)

// Injectors from wire.go:

func InitApp(cfg *config.Config) (*http.Server, func(), error) {
	db := repository.NewMysql(cfg)
	data, cleanup, err := repository.NewData(db)
	if err != nil {
		return nil, nil, err
	}
	testRepo := repository.NewTestRepo(data)
	testUseCase := service.NewTestUseCase(testRepo)
	demoRepo := repository.NewDemoRepo(data)
	demoUseCase := service.NewDemoUseCase(demoRepo)
	controllerController := controller.NewController(testUseCase, demoUseCase)
	engine := router.InitRouter(controllerController)
	server := initServer(cfg, engine)
	return server, func() {
		cleanup()
	}, nil
}
