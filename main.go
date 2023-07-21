package main

import (

	"my-go-micro/handler"
	pb "my-go-micro/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

)

var (
	service = "my-go-micro"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterMyGoMicroHandler(srv.Server(), new(handler.MyGoMicro)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
