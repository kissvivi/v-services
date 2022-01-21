package controller

import (
	"github.com/google/wire"
	"v-template-02/internal/service"
)

var ProviderSet = wire.NewSet(NewController)

type Controller struct {
	testService *service.TestUseCase
	demoService *service.DemoUseCase
}

func NewController(testService *service.TestUseCase, demoService *service.DemoUseCase) *Controller {
	return &Controller{testService: testService, demoService: demoService}
}
