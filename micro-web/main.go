package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/consul/v2"
	"log"
	"micro-web/route"
)

func main() {
	router := gin.Default()
	route.InitRouter(router)
	newRegistry := consul.NewRegistry(
		registry.Addrs("192.168.13.130"),
	)
	service := web.NewService(
		web.Name("micro-web"),
		web.Version("latest"),
		web.Handler(router),
		web.Registry(newRegistry),
		web.Address(":8081"),
	)
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
