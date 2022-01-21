package repository

import (
	"fmt"
	"v-template-01/api/http/request"
	"v-template-01/internal/service"
)

var _ service.DemoRepo = (*demoRepo)(nil)

type demoRepo struct {
	data *Data
}

func NewDemoRepo(data *Data) service.DemoRepo {
	return &demoRepo{data: data}
}

func (d demoRepo) DemoOk(request *request.TestRequest) string {
	fmt.Println("demo ok !!!!!!")

	return request.Name
}
