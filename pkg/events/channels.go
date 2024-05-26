package events

import (
	"GoExpertPostGrad-Orders-Challenge/configs"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQService struct {
	Channels map[string]*amqp.Channel
}

func NewRabbitMQService(config *configs.Envs, channelNames ...string) *RabbitMQService {
	channels := make(map[string]*amqp.Channel)
	for _, name := range channelNames {
		ch := getRabbitMQChannel(config)
		channels[name] = ch
		declareQueueAndBind(ch, name)
	}
	return &RabbitMQService{Channels: channels}
}

func getRabbitMQChannel(config *configs.Envs) *amqp.Channel {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.RabbitMQUser, config.RabbitMQPassword, config.RabbitMQHost, config.RabbitMQPort)
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel: %v", err)
	}
	return ch
}

func declareQueueAndBind(ch *amqp.Channel, queueName string) {
	_, err := ch.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	// Bind the queue to the exchange
	err = ch.QueueBind(
		queueName,                    // queue name
		fmt.Sprintf(queueName+"key"), // routing key
		"amq.direct",                 // exchange
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind queue: %v", err)
	}
}

func (r *RabbitMQService) Close() {
	for _, channel := range r.Channels {
		channel.Close()
	}
}
