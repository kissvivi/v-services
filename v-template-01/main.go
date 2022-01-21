package main

import "v-template-01/cmd"

// @title go-starter restful API
// @version 1.0.0
// @description go-starter
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://localhost:8080
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	//TODO 整理文件  优化结构 去除冗余代码

	cmd.Execute()
}
