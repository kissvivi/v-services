package service

import (
	"v-template-01/api/http/request"
	"v-template-01/internal/model"
)

type TestRepo interface {
	TestOk(request *request.TestRequest) string
	CreateTest(request *request.TestRequest) (model.Test, error)
}

type TestUseCase struct {
	testRepo TestRepo
}

func NewTestUseCase(testRepo TestRepo) *TestUseCase {
	return &TestUseCase{testRepo: testRepo}
}

//var _ TestRepo = (*TestUseCase)(nil)

func (test *TestUseCase) TestOk(request *request.TestRequest) string {
	return test.testRepo.TestOk(request)
}

func (test *TestUseCase) CreateTest(request *request.TestRequest) (model.Test, error) {
	return test.testRepo.CreateTest(request)
}
