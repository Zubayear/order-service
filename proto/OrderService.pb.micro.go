// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/OrderService.proto

package OrderService

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

// Api Endpoints for OrderService service

func NewOrderServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OrderService service

type OrderService interface {
	GetOrdersForUser(ctx context.Context, in *OrdersForUserRequest, opts ...client.CallOption) (*OrdersForUserResponse, error)
	AddOrder(ctx context.Context, in *AddOrderRequest, opts ...client.CallOption) (*AddOrderResponse, error)
	GetOrderById(ctx context.Context, in *OrderByIdRequest, opts ...client.CallOption) (*OrderByIdResponse, error)
	UpdateOrderPaymentStatus(ctx context.Context, in *OrderPaymentStatusRequest, opts ...client.CallOption) (*OrderPaymentStatusResponse, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) GetOrdersForUser(ctx context.Context, in *OrdersForUserRequest, opts ...client.CallOption) (*OrdersForUserResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.GetOrdersForUser", in)
	out := new(OrdersForUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) AddOrder(ctx context.Context, in *AddOrderRequest, opts ...client.CallOption) (*AddOrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.AddOrder", in)
	out := new(AddOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) GetOrderById(ctx context.Context, in *OrderByIdRequest, opts ...client.CallOption) (*OrderByIdResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.GetOrderById", in)
	out := new(OrderByIdResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) UpdateOrderPaymentStatus(ctx context.Context, in *OrderPaymentStatusRequest, opts ...client.CallOption) (*OrderPaymentStatusResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.UpdateOrderPaymentStatus", in)
	out := new(OrderPaymentStatusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderService service

type OrderServiceHandler interface {
	GetOrdersForUser(context.Context, *OrdersForUserRequest, *OrdersForUserResponse) error
	AddOrder(context.Context, *AddOrderRequest, *AddOrderResponse) error
	GetOrderById(context.Context, *OrderByIdRequest, *OrderByIdResponse) error
	UpdateOrderPaymentStatus(context.Context, *OrderPaymentStatusRequest, *OrderPaymentStatusResponse) error
}

func RegisterOrderServiceHandler(s server.Server, hdlr OrderServiceHandler, opts ...server.HandlerOption) error {
	type orderService interface {
		GetOrdersForUser(ctx context.Context, in *OrdersForUserRequest, out *OrdersForUserResponse) error
		AddOrder(ctx context.Context, in *AddOrderRequest, out *AddOrderResponse) error
		GetOrderById(ctx context.Context, in *OrderByIdRequest, out *OrderByIdResponse) error
		UpdateOrderPaymentStatus(ctx context.Context, in *OrderPaymentStatusRequest, out *OrderPaymentStatusResponse) error
	}
	type OrderService struct {
		orderService
	}
	h := &orderServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OrderService{h}, opts...))
}

type orderServiceHandler struct {
	OrderServiceHandler
}

func (h *orderServiceHandler) GetOrdersForUser(ctx context.Context, in *OrdersForUserRequest, out *OrdersForUserResponse) error {
	return h.OrderServiceHandler.GetOrdersForUser(ctx, in, out)
}

func (h *orderServiceHandler) AddOrder(ctx context.Context, in *AddOrderRequest, out *AddOrderResponse) error {
	return h.OrderServiceHandler.AddOrder(ctx, in, out)
}

func (h *orderServiceHandler) GetOrderById(ctx context.Context, in *OrderByIdRequest, out *OrderByIdResponse) error {
	return h.OrderServiceHandler.GetOrderById(ctx, in, out)
}

func (h *orderServiceHandler) UpdateOrderPaymentStatus(ctx context.Context, in *OrderPaymentStatusRequest, out *OrderPaymentStatusResponse) error {
	return h.OrderServiceHandler.UpdateOrderPaymentStatus(ctx, in, out)
}
