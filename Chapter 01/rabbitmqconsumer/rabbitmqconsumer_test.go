package rabbitmqconsumer

import (
	"context"
	"testing"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
)

func TestReceiveMessage(t *testing.T) {
	consumer, connection := initTest()

	go func() {
		time.Sleep(2 * time.Second)
		consumer.channel.Close()
		connection.Close()
	}()

	err := consumer.ReceiveMessage()
	assert.Equal(t, nil, err)
}

func initTest() (*RabbitMQConsumer, *amqp.Connection) {
	// Init consumer
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	channel, _ := conn.Channel()
	queue, _ := channel.QueueDeclare("test", false, false, false, false, nil)

	// Send test message
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Test message"),
	}

	channel.PublishWithContext(context.Background(), "", queue.Name, false, false, message)

	// Return consumer
	return &RabbitMQConsumer{
		channel: channel,
		queue:   &queue,
	}, conn
}
