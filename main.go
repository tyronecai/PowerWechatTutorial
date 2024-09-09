package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"power-wechat-tutorial/config"
	"power-wechat-tutorial/routes"
	"power-wechat-tutorial/services"
)

var Host = ""
var Port = "8888"

// @title           PowerWechat API Docs
// @version         1.0.1
// @description     这是一个开源的使用教程，包含PowerWechat的大部分接口调试代码.
// @termsOfService  http://artisan-cloud.com/terms/

// @contact.name   ArtisanCloud Support
// @contact.url    https://powerwechat.artisan-cloud.com/zh/start/qa.html
// @contact.email  matrix-x@artisan-cloud.como

// @license.name  MIT 2.0
// @license.url   https://github.com/ArtisanCloud/PowerWeChat?tab=MIT-1-ov-file#readme

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	conf := config.Get()

	var err error
	services.PaymentApp, err = services.NewWXPaymentApp(conf)
	if err != nil || services.PaymentApp == nil {
		panic(err)
	}

	services.MiniProgramApp, err = services.NewMiniMiniProgramService(conf)
	if err != nil || services.MiniProgramApp == nil {
		panic(err)
	}

	services.OfficialAccountApp, err = services.NewOfficialAccountAppService(conf)
	if err != nil || services.OfficialAccountApp == nil {
		panic(err)
	}

	services.WeComApp, err = services.NewWeComService(conf)
	if err != nil || services.WeComApp == nil {
		panic(err)
	}

	services.WeComContactApp, err = services.NewWeComContactService(conf)
	if err != nil || services.WeComContactApp == nil {
		panic(err)
	}

	services.OpenPlatformApp, err = services.NewOpenPlatformAppService(conf)
	if err != nil || services.OpenPlatformApp == nil {
		panic(err)
	}

	r := gin.Default()

	// Initialize the routes
	routes.InitializeRoutes(r)

	log.Fatalln(r.Run(Host + ":" + Port))

}
