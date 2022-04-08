package repository

import (
	"OrderService/ent"
	"OrderService/ent/order"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type IOrderRepository interface {
	SaveOrder(ctx context.Context, order *ent.Order) (*ent.Order, error)
	FetchOrdersForUser(ctx context.Context, userId uuid.UUID) ([]*ent.Order, error)
	FetchOrderById(ctx context.Context, orderId uuid.UUID) (*ent.Order, error)
	ModifyOrderPaymentStatus(ctx context.Context, orderId uuid.UUID, paid bool) error
}

type OrderRepository struct {
	client *ent.Client
}

func (o *OrderRepository) SaveOrder(ctx context.Context, order *ent.Order) (*ent.Order, error) {
	savedOrder, err := o.client.Order.Create().
		SetOrderPaid(order.OrderPaid).
		SetOrderTotal(order.OrderTotal).
		SetUserID(order.UserID).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed saving order: %w", err)
	}
	return savedOrder, nil
}

func (o *OrderRepository) FetchOrdersForUser(ctx context.Context, userId uuid.UUID) ([]*ent.Order, error) {
	orders, err := o.client.Order.Query().Where(order.UserID(userId)).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed getting orders: %w", err)
	}
	return orders, nil
}

func (o *OrderRepository) FetchOrderById(ctx context.Context, orderId uuid.UUID) (*ent.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o *OrderRepository) ModifyOrderPaymentStatus(ctx context.Context, orderId uuid.UUID, paid bool) error {
	err := o.client.Order.UpdateOneID(orderId).SetOrderPaid(paid).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func OrderRepositoryProvider(client *ent.Client) (*OrderRepository, error) {
	return &OrderRepository{client: client}, nil
}

func NewEventRepository() IOrderRepository {
	return &OrderRepository{}
}
