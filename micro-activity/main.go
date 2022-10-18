package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
	"micro-acticity/control"
	"micro-acticity/proto"
)

func main() {
	newRegistry := consul.NewRegistry(
		registry.Addrs("192.168.13.130:8500"),
	)
	service := micro.NewService(
		micro.Name("micro-activity"),
		micro.Version("latest"),
		micro.Registry(newRegistry),
	)
	proto.RegisterActivityHandler(service.Server(), new(control.Activity))
	service.Init()
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
