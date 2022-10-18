package choose

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"lottery/micro-choose/proto"
)

var ChooseService proto.ChooseService

func init() {
	registry := consul.NewRegistry()
	service := micro.NewService(
		micro.Name("micro-choose"),
		micro.Registry(registry),
	)
	ChooseService = proto.NewChooseService("micro-choose", service.Client())
}
