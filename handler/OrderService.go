package handler

import (
	"OrderService/ent"
	pb "OrderService/proto"
	"OrderService/repository"
	"context"
	"fmt"
	"github.com/google/uuid"

	"go-micro.dev/v4/logger"
)

type OrderService struct{ repo repository.IOrderRepository }

func (o *OrderService) GetOrdersForUser(ctx context.Context, req *pb.OrdersForUserRequest, rsp *pb.OrdersForUserResponse) error {
	logger.Infof("Received OrderService.GetOrdersForUser request: %v", req)
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return fmt.Errorf("failed parsting user id: %w", err)
	}
	ordersForUser, err := o.repo.FetchOrdersForUser(ctx, userId)
	if err != nil {
		return err
	}
	var ordersToReturn []*pb.Order
	for _, order := range ordersForUser {
		mappedOrder := &pb.Order{}
		mappedOrder.UserId = order.UserID.String()
		mappedOrder.Id = order.ID.String()
		mappedOrder.OrderTotal = float32(order.OrderTotal)
		mappedOrder.OrderPaid = order.OrderPaid
		mappedOrder.OrderPlaced = order.OrderPlaced.Unix()
		ordersToReturn = append(ordersToReturn, mappedOrder)
	}
	rsp.Order = ordersToReturn
	return nil
}

func (o *OrderService) AddOrder(ctx context.Context, req *pb.AddOrderRequest, rsp *pb.AddOrderResponse) error {
	logger.Infof("Received OrderService.AddOrder request: %v", req)
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return fmt.Errorf("failed parsing user id: %w", err)
	}
	orderToSave := &ent.Order{
		UserID:     userId,
		OrderTotal: float64(req.OrderTotal),
		OrderPaid:  req.OrderPaid,
	}
	orderFromRepo, err := o.repo.SaveOrder(ctx, orderToSave)
	if err != nil {
		return err
	}
	orderToReturn := &pb.Order{
		Id:          orderFromRepo.ID.String(),
		UserId:      orderFromRepo.UserID.String(),
		OrderTotal:  float32(orderFromRepo.OrderTotal),
		OrderPlaced: orderFromRepo.OrderPlaced.Unix(),
		OrderPaid:   orderFromRepo.OrderPaid,
	}
	rsp.Order = orderToReturn
	return nil
}
func (o *OrderService) GetOrderById(ctx context.Context, req *pb.OrderByIdRequest, rsp *pb.OrderByIdResponse) error {
	logger.Infof("Received OrderService.GetOrderById request: %v", req)
	orderId, err := uuid.Parse(req.OrderId)
	if err != nil {
		return fmt.Errorf("failed parsting order id: %w", err)
	}
	orderFromRepo, err := o.repo.FetchOrderById(ctx, orderId)
	if err != nil {
		return err
	}
	orderToReturn := &pb.Order{
		Id:          orderFromRepo.ID.String(),
		UserId:      orderFromRepo.UserID.String(),
		OrderTotal:  float32(orderFromRepo.OrderTotal),
		OrderPlaced: orderFromRepo.OrderPlaced.Unix(),
		OrderPaid:   orderFromRepo.OrderPaid,
	}
	rsp.Order = orderToReturn
	return nil
}

func (o *OrderService) UpdateOrderPaymentStatus(ctx context.Context, req *pb.OrderPaymentStatusRequest, rsp *pb.OrderPaymentStatusResponse) error {
	logger.Infof("Received OrderService.UpdateOrderPaymentStatus request: %v", req)
	return nil
}

func NewOrderProvider(repo repository.IOrderRepository) (*OrderService, error) {
	return &OrderService{repo: repo}, nil
}
