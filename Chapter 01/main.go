package main

import (
	"log"
	"receiver-rabbitmq/rabbitmqconsumer"
)

func main() {
	log.Println("Starting Receiver Service")
	consumer := rabbitmqconsumer.InitRabbitMQConsumer()

	consumer.ReceiveMessage()
}
