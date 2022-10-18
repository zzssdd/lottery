package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
	"micro-user/control"
	"micro-user/proto/admin"
	"micro-user/proto/user"
)

func main() {
	newRegistry := consul.NewRegistry(
		registry.Addrs("192.168.13.130:8500"),
	)
	service := micro.NewService(
		micro.Name("micro-user"),
		micro.Version("latest"),
		micro.Registry(newRegistry),
	)
	service.Init()
	user.RegisterUserHandler(service.Server(), new(control.User))
	admin.RegisterAdminHandler(service.Server(), new(control.Admin))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
