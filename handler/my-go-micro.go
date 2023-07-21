package handler

import (
	"context"

	"go-micro.dev/v4/logger"

	pb "my-go-micro/proto"
)

type MyGoMicro struct{}

func (e *MyGoMicro) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received MyGoMicro.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}
