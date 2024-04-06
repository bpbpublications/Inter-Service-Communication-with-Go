package main

import (
	"context"
	"fmt"
	"log"
	"sender-rabbitmq/rabbitmqproducer"
)

func main() {
	log.Println("Starting Sender Service")
	producer := rabbitmqproducer.InitRabbitMQProducer()

	defer producer.Close()

	for i := 0; i < 10; i++ {
		err := producer.SendMessage(context.Background(), "Hello World!!!")
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	fmt.Println("All messages are sent !!!")
}
