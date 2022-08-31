// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: lang/go-micro/proto/cap/hello.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Cap service

func NewCapEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Cap service

type CapService interface {
	SayHello(ctx context.Context, in *SayRequest, opts ...client.CallOption) (*SayResponse, error)
}

type capService struct {
	c    client.Client
	name string
}

func NewCapService(name string, c client.Client) CapService {
	return &capService{
		c:    c,
		name: name,
	}
}

func (c *capService) SayHello(ctx context.Context, in *SayRequest, opts ...client.CallOption) (*SayResponse, error) {
	req := c.c.NewRequest(c.name, "Cap.SayHello", in)
	out := new(SayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cap service

type CapHandler interface {
	SayHello(context.Context, *SayRequest, *SayResponse) error
}

func RegisterCapHandler(s server.Server, hdlr CapHandler, opts ...server.HandlerOption) error {
	type cap interface {
		SayHello(ctx context.Context, in *SayRequest, out *SayResponse) error
	}
	type Cap struct {
		cap
	}
	h := &capHandler{hdlr}
	return s.Handle(s.NewHandler(&Cap{h}, opts...))
}

type capHandler struct {
	CapHandler
}

func (h *capHandler) SayHello(ctx context.Context, in *SayRequest, out *SayResponse) error {
	return h.CapHandler.SayHello(ctx, in, out)
}
