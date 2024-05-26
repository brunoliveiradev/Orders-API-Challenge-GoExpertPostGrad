package handler

import (
	"GoExpertPostGrad-Orders-Challenge/pkg/cache"
	"GoExpertPostGrad-Orders-Challenge/pkg/events"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type OrderCreatedHandler struct {
	RabbitMQService *events.RabbitMQService
	Cache           cache.Cache
}

func NewOrderCreatedHandler(rabbitMQService *events.RabbitMQService, cache cache.Cache) *OrderCreatedHandler {
	return &OrderCreatedHandler{RabbitMQService: rabbitMQService, Cache: cache}
}

func (h *OrderCreatedHandler) Handle(event events.EventInterface) error {
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

	// Clear cache when a new order is created
	h.Cache.Delete("orders")
	return nil
}

func (h *OrderCreatedHandler) GetEventName() string {
	return "OrderCreated"
}
