package consumer

import (
	"OrderService/ent"
	"OrderService/messages"
	"OrderService/producer"
	"OrderService/repository"
	"context"
	"encoding/json"
	"fmt"
	"go-micro.dev/v4/broker"
	log "go-micro.dev/v4/logger"
)

var (
	topicCheckoutMessage    = "shoppingBasket.topic.checkoutMessage"
	topicOrderPayment       = "order.topic.orderPayment"
	topicOrderPaymentUpdate = "order.topic.orderPaymentUpdate"
)

type MessageConsumer struct {
	repo repository.IOrderRepository
}

func MessageConsumerProvider(repo repository.IOrderRepository) (*MessageConsumer, error) {
	return &MessageConsumer{repo: repo}, nil
}

func (mc *MessageConsumer) ConsumeCheckoutMessage() {
	_, err := broker.Subscribe(topicCheckoutMessage, func(event broker.Event) error {
		var consumedMessage messages.BasketCheckoutMessage
		err := json.Unmarshal(event.Message().Body, &consumedMessage)
		if err != nil {
			return fmt.Errorf("failed unmarshalling message: %w", err)
		}
		log.Infof("consumed message consumedMessage: %v", consumedMessage)
		var orderToSave ent.Order
		orderToSave.UserID = consumedMessage.UserId
		orderToSave.OrderTotal = consumedMessage.BasketTotal
		orderFromRepo, err := mc.repo.SaveOrder(context.Background(), &orderToSave)
		if err != nil {
			return err
		}
		orderPaymentRequestMessage := messages.NewOrderPaymentRequestMessage(
			orderFromRepo.ID,
			consumedMessage.BasketTotal,
			consumedMessage.CardNumber,
			consumedMessage.CardName,
			consumedMessage.CardExpiration)
		err = producer.PublishMessage(*orderPaymentRequestMessage, topicOrderPayment)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
}

func (mc *MessageConsumer) ConsumeOrderPaymentUpdateMessage() {
	_, err := broker.Subscribe(topicOrderPaymentUpdate, func(event broker.Event) error {
		var orderPaymentUpdate messages.OrderPaymentUpdateMessage
		err := json.Unmarshal(event.Message().Body, &orderPaymentUpdate)
		if err != nil {
			return err
		}
		log.Infof("consumed message orderPaymentUpdate: %v", orderPaymentUpdate)
		err = mc.repo.ModifyOrderPaymentStatus(context.Background(), orderPaymentUpdate.OrderId, orderPaymentUpdate.PaymentStatus)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
}
