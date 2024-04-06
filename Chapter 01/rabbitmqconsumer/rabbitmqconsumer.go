package rabbitmqconsumer

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConsumer struct {
	channel *amqp.Channel
	queue   *amqp.Queue
}

func InitRabbitMQConsumer() *RabbitMQConsumer {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("RabbitMQ consumer error: %v", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("RabbitMQ consumer error: %v", err)
	}

	queue, err := channel.QueueDeclare("messages", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("RabbitMQ consumer error: %v", err)
	}

	go func() {
		// Listen to operating system's interrupt signal
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
		<-interrupt

		// Gracefully shut down the consumer when it happens
		conn.Close()
		channel.Close()
	}()

	return &RabbitMQConsumer{
		channel: channel,
		queue:   &queue,
	}
}

func (rc *RabbitMQConsumer) ReceiveMessage() error {
	messages, err := rc.channel.Consume(rc.queue.Name, "consumer", true, false, false, true, nil)
	if err != nil {
		return err
	}

	for message := range messages {
		fmt.Println(string(message.Body))
	}

	return nil
}
