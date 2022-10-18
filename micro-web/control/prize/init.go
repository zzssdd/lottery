package prize

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"lottery/micro-prize/proto"
)

var PrizeService proto.PrizeService

func init() {
	registry := consul.NewRegistry()
	service := micro.NewService(
		micro.Name("micro-prize"),
		micro.Registry(registry),
	)
	PrizeService = proto.NewPrizeService("micro-prize", service.Client())
}
