package order

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"lottery/micro-order/proto"
)

var OrderService proto.OrderService

func init() {
	registry := consul.NewRegistry()
	service := micro.NewService(
		micro.Name("micro-order"),
		micro.Registry(registry),
	)
	OrderService = proto.NewOrderService("micro-order", service.Client())
}
