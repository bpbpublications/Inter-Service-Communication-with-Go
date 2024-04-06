package rabbitmqproducer

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQProducer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      *amqp.Queue
}

func InitRabbitMQProducer() *RabbitMQProducer {
	connection, err := amqp.Dial(
		"amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("RabbitMQ producer error: %v", err)
	}

	channel, err := connection.Channel()
	if err != nil {
		log.Fatalf("RabbitMQ producer error: %v", err)
	}

	queue, err := channel.QueueDeclare("messages", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("RabbitMQ producer error: %v", err)
	}

	return &RabbitMQProducer{
		connection: connection,
		channel:    channel,
		queue:      &queue,
	}
}

func (rp *RabbitMQProducer) SendMessage(ctx context.Context, content string) error {
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(content),
	}

	return rp.channel.PublishWithContext(ctx, "", rp.queue.Name, false, false, message)
}

func (rp *RabbitMQProducer) Close() {
	rp.connection.Close()
	rp.channel.Close()
}
