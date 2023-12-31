// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/my-go-micro.proto

package mygomicro

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for MyGoMicro service

func NewMyGoMicroEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MyGoMicro service

type MyGoMicroService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
}

type myGoMicroService struct {
	c    client.Client
	name string
}

func NewMyGoMicroService(name string, c client.Client) MyGoMicroService {
	return &myGoMicroService{
		c:    c,
		name: name,
	}
}

func (c *myGoMicroService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "MyGoMicro.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MyGoMicro service

type MyGoMicroHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
}

func RegisterMyGoMicroHandler(s server.Server, hdlr MyGoMicroHandler, opts ...server.HandlerOption) error {
	type myGoMicro interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
	}
	type MyGoMicro struct {
		myGoMicro
	}
	h := &myGoMicroHandler{hdlr}
	return s.Handle(s.NewHandler(&MyGoMicro{h}, opts...))
}

type myGoMicroHandler struct {
	MyGoMicroHandler
}

func (h *myGoMicroHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.MyGoMicroHandler.Call(ctx, in, out)
}
