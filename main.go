package main

import (
	"OrderService/di"
	pb "OrderService/proto"
	"flag"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/cmd"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "orderservice"
	version = "latest"
	port    = flag.String("port", ":60010", "order port")
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Address(*port),
	)
	srv.Init()

	orderService, err := di.DependencyProvider()
	if err != nil {
		return
	}
	// Register handler
	err = pb.RegisterOrderServiceHandler(srv.Server(), orderService)
	if err != nil {
		return
	}

	// Broker Section
	err = cmd.Init()
	if err != nil {
		return
	}

	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	brokerDependency, err := di.BrokerDependency()
	if err != nil {
		return
	}

	go brokerDependency.ConsumeCheckoutMessage()
	go brokerDependency.ConsumeOrderPaymentUpdateMessage()

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
