package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
	"micro-order/control"
	"micro-order/proto"
)

func main() {
	newRegistry := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	service := micro.NewService(
		micro.Name("micro-order"),
		micro.Version("latest"),
		micro.Registry(newRegistry),
	)
	service.Init()
	proto.RegisterOrderHandler(service.Server(), new(control.Order))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
