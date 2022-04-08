//go:build wireinject
// +build wireinject

package di

import (
	"OrderService/consumer"
	"OrderService/external"
	"OrderService/handler"
	"OrderService/repository"
	"github.com/google/wire"
)

func DependencyProvider() (*handler.OrderService, error) {
	wire.Build(
		repository.ClientProvider,
		external.HostProvider,
		repository.OrderRepositoryProvider,
		handler.NewOrderProvider,
		wire.Bind(new(repository.IOrderRepository), new(*repository.OrderRepository)))
	return &handler.OrderService{}, nil
}

func BrokerDependency() (*consumer.MessageConsumer, error) {
	wire.Build(
		repository.ClientProvider,
		external.HostProvider,
		repository.OrderRepositoryProvider,
		consumer.MessageConsumerProvider,
		wire.Bind(new(repository.IOrderRepository), new(*repository.OrderRepository)))
	return &consumer.MessageConsumer{}, nil
}
