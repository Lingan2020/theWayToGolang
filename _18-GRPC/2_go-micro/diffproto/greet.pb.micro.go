// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: greet.proto

package greetpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for GreetService service

type GreetService interface {
	// Unary
	Greet(ctx context.Context, in *GreetRequest, opts ...client.CallOption) (*GreetResponse, error)
	// Server Streaming
	GreetManyTimes(ctx context.Context, in *GreetManyTimesRequest, opts ...client.CallOption) (GreetService_GreetManyTimesService, error)
	// Client Streaming
	LongGreet(ctx context.Context, opts ...client.CallOption) (GreetService_LongGreetService, error)
	// BiDi Streaming
	GreetEveryone(ctx context.Context, opts ...client.CallOption) (GreetService_GreetEveryoneService, error)
}

type greetService struct {
	c    client.Client
	name string
}

func NewGreetService(name string, c client.Client) GreetService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "greet"
	}
	return &greetService{
		c:    c,
		name: name,
	}
}

func (c *greetService) Greet(ctx context.Context, in *GreetRequest, opts ...client.CallOption) (*GreetResponse, error) {
	req := c.c.NewRequest(c.name, "GreetService.Greet", in)
	out := new(GreetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetService) GreetManyTimes(ctx context.Context, in *GreetManyTimesRequest, opts ...client.CallOption) (GreetService_GreetManyTimesService, error) {
	req := c.c.NewRequest(c.name, "GreetService.GreetManyTimes", &GreetManyTimesRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &greetServiceGreetManyTimes{stream}, nil
}

type GreetService_GreetManyTimesService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*GreetManytimesResponse, error)
}

type greetServiceGreetManyTimes struct {
	stream client.Stream
}

func (x *greetServiceGreetManyTimes) Close() error {
	return x.stream.Close()
}

func (x *greetServiceGreetManyTimes) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *greetServiceGreetManyTimes) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *greetServiceGreetManyTimes) Recv() (*GreetManytimesResponse, error) {
	m := new(GreetManytimesResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetService) LongGreet(ctx context.Context, opts ...client.CallOption) (GreetService_LongGreetService, error) {
	req := c.c.NewRequest(c.name, "GreetService.LongGreet", &LongGreetRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &greetServiceLongGreet{stream}, nil
}

type GreetService_LongGreetService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*LongGreetRequest) error
}

type greetServiceLongGreet struct {
	stream client.Stream
}

func (x *greetServiceLongGreet) Close() error {
	return x.stream.Close()
}

func (x *greetServiceLongGreet) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *greetServiceLongGreet) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *greetServiceLongGreet) Send(m *LongGreetRequest) error {
	return x.stream.Send(m)
}

func (c *greetService) GreetEveryone(ctx context.Context, opts ...client.CallOption) (GreetService_GreetEveryoneService, error) {
	req := c.c.NewRequest(c.name, "GreetService.GreetEveryone", &GreetEveryoneRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &greetServiceGreetEveryone{stream}, nil
}

type GreetService_GreetEveryoneService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*GreetEveryoneRequest) error
	Recv() (*GreetEveryoneResponse, error)
}

type greetServiceGreetEveryone struct {
	stream client.Stream
}

func (x *greetServiceGreetEveryone) Close() error {
	return x.stream.Close()
}

func (x *greetServiceGreetEveryone) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *greetServiceGreetEveryone) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *greetServiceGreetEveryone) Send(m *GreetEveryoneRequest) error {
	return x.stream.Send(m)
}

func (x *greetServiceGreetEveryone) Recv() (*GreetEveryoneResponse, error) {
	m := new(GreetEveryoneResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GreetService service

type GreetServiceHandler interface {
	// Unary
	Greet(context.Context, *GreetRequest, *GreetResponse) error
	// Server Streaming
	GreetManyTimes(context.Context, *GreetManyTimesRequest, GreetService_GreetManyTimesStream) error
	// Client Streaming
	LongGreet(context.Context, GreetService_LongGreetStream) error
	// BiDi Streaming
	GreetEveryone(context.Context, GreetService_GreetEveryoneStream) error
}

func RegisterGreetServiceHandler(s server.Server, hdlr GreetServiceHandler, opts ...server.HandlerOption) error {
	type greetService interface {
		Greet(ctx context.Context, in *GreetRequest, out *GreetResponse) error
		GreetManyTimes(ctx context.Context, stream server.Stream) error
		LongGreet(ctx context.Context, stream server.Stream) error
		GreetEveryone(ctx context.Context, stream server.Stream) error
	}
	type GreetService struct {
		greetService
	}
	h := &greetServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&GreetService{h}, opts...))
}

type greetServiceHandler struct {
	GreetServiceHandler
}

func (h *greetServiceHandler) Greet(ctx context.Context, in *GreetRequest, out *GreetResponse) error {
	return h.GreetServiceHandler.Greet(ctx, in, out)
}

func (h *greetServiceHandler) GreetManyTimes(ctx context.Context, stream server.Stream) error {
	m := new(GreetManyTimesRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.GreetServiceHandler.GreetManyTimes(ctx, m, &greetServiceGreetManyTimesStream{stream})
}

type GreetService_GreetManyTimesStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*GreetManytimesResponse) error
}

type greetServiceGreetManyTimesStream struct {
	stream server.Stream
}

func (x *greetServiceGreetManyTimesStream) Close() error {
	return x.stream.Close()
}

func (x *greetServiceGreetManyTimesStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *greetServiceGreetManyTimesStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *greetServiceGreetManyTimesStream) Send(m *GreetManytimesResponse) error {
	return x.stream.Send(m)
}

func (h *greetServiceHandler) LongGreet(ctx context.Context, stream server.Stream) error {
	return h.GreetServiceHandler.LongGreet(ctx, &greetServiceLongGreetStream{stream})
}

type GreetService_LongGreetStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*LongGreetRequest, error)
}

type greetServiceLongGreetStream struct {
	stream server.Stream
}

func (x *greetServiceLongGreetStream) Close() error {
	return x.stream.Close()
}

func (x *greetServiceLongGreetStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *greetServiceLongGreetStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *greetServiceLongGreetStream) Recv() (*LongGreetRequest, error) {
	m := new(LongGreetRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *greetServiceHandler) GreetEveryone(ctx context.Context, stream server.Stream) error {
	return h.GreetServiceHandler.GreetEveryone(ctx, &greetServiceGreetEveryoneStream{stream})
}

type GreetService_GreetEveryoneStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*GreetEveryoneResponse) error
	Recv() (*GreetEveryoneRequest, error)
}

type greetServiceGreetEveryoneStream struct {
	stream server.Stream
}

func (x *greetServiceGreetEveryoneStream) Close() error {
	return x.stream.Close()
}

func (x *greetServiceGreetEveryoneStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *greetServiceGreetEveryoneStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *greetServiceGreetEveryoneStream) Send(m *GreetEveryoneResponse) error {
	return x.stream.Send(m)
}

func (x *greetServiceGreetEveryoneStream) Recv() (*GreetEveryoneRequest, error) {
	m := new(GreetEveryoneRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
