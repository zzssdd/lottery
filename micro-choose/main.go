package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
	"micro-choose/control"
	"micro-choose/proto"
)

func main() {
	newRegistry := consul.NewRegistry(
		registry.Addrs("192.168.13.130:8500"),
	)
	service := micro.NewService(
		micro.Name("micro-choose"),
		micro.Version("latest"),
		micro.Registry(newRegistry),
	)
	service.Init()
	proto.RegisterChooseHandler(service.Server(), new(control.Choose))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
