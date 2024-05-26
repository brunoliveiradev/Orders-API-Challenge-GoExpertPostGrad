package handler

import (
	"GoExpertPostGrad-Orders-Challenge/pkg/events"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type OrderListingHandler struct {
	RabbitMQService *events.RabbitMQService
}

func NewOrderListingHandler(rabbitMQService *events.RabbitMQService) *OrderListingHandler {
	return &OrderListingHandler{RabbitMQService: rabbitMQService}
}

func (h *OrderListingHandler) Handle(event events.EventInterface) error {
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msg := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	channel, ok := h.RabbitMQService.Channels["orders"]
	if !ok {
		log.Println("Error: channel 'orders' not found")
		return fmt.Errorf("channel 'orders' not found")
	}

	err := channel.Publish(
		"amq.direct", // exchange
		"order_key",  // routing key
		false,        // mandatory
		false,        // immediate
		msg,          // message to publish
	)
	if err != nil {
		log.Println("Error publishing event:", err)
		return err
	}
	log.Println("Event published", event.GetName(), ":", string(jsonOutput))

	return nil
}

func (h *OrderListingHandler) GetEventName() string {
	return "OrderListing"
}
