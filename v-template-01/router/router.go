package router

import (
	"github.com/gin-gonic/gin"
	"v-template-01/internal/controller"
)

func routers(r *gin.Engine, controller *controller.Controller) {
	g := r.Group("/test")
	{
		g.GET("/ok", controller.TestOk)
		g.GET("/demo/ok", controller.DemoOk)
		g.POST("/create", controller.CreateTest)
	}
}
