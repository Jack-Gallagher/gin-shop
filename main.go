package main

import (
	"ginShop/controller"
	"ginShop/tool"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}
	app := gin.Default()

	registerHello(app)
	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

func registerHello(router *gin.Engine) {
	new(controller.HelloController).Router(router)
}
