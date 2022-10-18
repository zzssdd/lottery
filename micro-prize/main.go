package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
	"micro-prize/control"
	"micro-prize/proto"
)

func main() {
	newRegistry := consul.NewRegistry(registry.Addrs("192.168.13.130:8500"))
	service := micro.NewService(
		micro.Name("micro-prize"),
		micro.Version("latest"),
		micro.Registry(newRegistry),
	)
	service.Init()
	proto.RegisterPrizeHandler(service.Server(), new(control.Prize))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
