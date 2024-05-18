package handler

import (
	"GoExpertPostGrad-Orders-Challenge/pkg/events"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

type OrderCreatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderCreatedHandler(rabbitMQChannel *amqp.Channel) *OrderCreatedHandler {
	_, err := rabbitMQChannel.QueueDeclare(
		"orders", // name of the queue
		true,     // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	// Bind the queue to the exchange
	err = rabbitMQChannel.QueueBind(
		"orders",     // queue name
		"order_key",  // routing key
		"amq.direct", // exchange
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind queue: %v", err)
	}

	return &OrderCreatedHandler{RabbitMQChannel: rabbitMQChannel}
}

func (h *OrderCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) error {
	defer wg.Done()

	jsonOutput, _ := json.Marshal(event.GetPayload())

	msg := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	err := h.RabbitMQChannel.Publish(
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
