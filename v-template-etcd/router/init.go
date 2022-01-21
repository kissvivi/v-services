package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"v-template-etcd/internal/controller"
	"v-template-etcd/pkg/registry/etcd"
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

	r.Use(httpService())
	routers(r, testController)

	return r
}

type TargetHost struct {
	Host    string
	IsHttps bool
	CAPath  string
}

func httpService() gin.HandlerFunc {

	return func(c *gin.Context) {

		balance := etcd.Balance{}
		service := balance.GetService()
		node, err := balance.SelectRandom(context.TODO(), service.Nodes)
		if err != nil {
			log.Printf("err%v", err)
		}
		host := fmt.Sprintf("%v:%v", node.IP, node.Port)

		log.Printf("转发服务名:%v，地址：%s", service.Name, host)

		u := &url.URL{}
		u.Scheme = "http"
		u.Host = host
		proxy := httputil.NewSingleHostReverseProxy(u)

		//重写出错回调
		proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
			log.Printf("http: proxy error: %v", err)
			ret := fmt.Sprintf("http proxy error %v", err)

			//写到body里
			rw.Write([]byte(ret))
		}
		log.Printf("%v", c.Request.URL)

		urrl := c.Request.RequestURI
		serviceNameUrl := fmt.Sprintf("/%s", service.Name)
		urrrl := strings.Replace(urrl, serviceNameUrl, "", 1)
		//fmt.Println(urrl)
		//fmt.Println(urrrl)
		c.Request.URL, err = url.Parse(urrrl)
		if err != nil {
			log.Printf("err%v", err)
		}
		proxy.ServeHTTP(c.Writer, c.Request)

		//fmt.Println("服务发现ing........")
		//fmt.Println(c.HandlerName())
	}
}
