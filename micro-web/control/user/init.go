package user

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"lottery/micro-user/proto/admin"
	"lottery/micro-user/proto/user"
)

var (
	UserService  user.UserService
	AdminService admin.AdminService
)

func init() {
	registry := consul.NewRegistry()
	service := micro.NewService(
		micro.Name("micro-user"),
		micro.Registry(registry),
	)
	UserService = user.NewUserService("micro-user", service.Client())
	AdminService = admin.NewAdminService("micro-user", service.Client())
}
