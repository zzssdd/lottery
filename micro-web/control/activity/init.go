package activity

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"lottery/micro-activity/proto"
)

var ActiveService proto.ActivityService

func init() {
	registry := consul.NewRegistry()
	service := micro.NewService(
		micro.Name("micro-activity"),
		micro.Registry(registry),
	)
	ActiveService = proto.NewActivityService("micro-activity", service.Client())
}
