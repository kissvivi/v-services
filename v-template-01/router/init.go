package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"v-template-01/internal/controller"
)

type engine = *gin.Engine

var InitRouterSet = wire.NewSet(InitRouter)

func InitRouter(testController *controller.Controller) engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode("debug")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong this is template-01"})
	})

	routers(r, testController)

	return r
}
