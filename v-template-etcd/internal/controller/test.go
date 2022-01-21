package controller

import (
	"github.com/gin-gonic/gin"
	"v-template-etcd/api/http/request"
	"v-template-etcd/pkg/app"
	"v-template-etcd/pkg/msg"
)

// service 依赖于repo
//var testRepo = repository.NewTestRepo()
//var testService = service.NewTestUseCase(testRepo)

//TODO 深入di依赖注入
//var t = initTestController()

func (t Controller) TestOk(c *gin.Context) {

	query := request.TestRequest{Id: 111, Name: "hello world!"}

	data := t.testService.TestOk(&query)

	app.OK(c, data, msg.OK)
}

func (t Controller) DemoOk(c *gin.Context) {

	query := request.TestRequest{Id: 111, Name: "hello world! --demo"}

	data := t.demoService.DemoOk(&query)

	app.OK(c, data, msg.OK)
}

func (t Controller) CreateTest(c *gin.Context) {
	query := request.TestRequest{Name: "hello world!"}
	data, err := t.testService.CreateTest(&query)
	if err != nil {
		app.Error(c, 500, msg.ERR)
	}
	app.OK(c, data, msg.OK)
}
