package main

import (
	"micro-choose/handler"
	pb "micro-choose/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("micro-choose"),
	)

	// Register handler
	pb.RegisterMicroChooseHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
