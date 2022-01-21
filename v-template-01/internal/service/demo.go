package service

import (
	"v-template-01/api/http/request"
)

type DemoRepo interface {
	DemoOk(request *request.TestRequest) string
}

type DemoUseCase struct {
	demoRepo DemoRepo
}

func NewDemoUseCase(demoRepo DemoRepo) *DemoUseCase {
	return &DemoUseCase{demoRepo: demoRepo}
}

func (test *DemoUseCase) DemoOk(request *request.TestRequest) string {
	return test.demoRepo.DemoOk(request)
}
