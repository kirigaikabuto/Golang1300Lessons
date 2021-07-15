package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
		return
	}
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
		return
	}
	queue, err := channel.QueueDeclare(
		"lesson18",
		false,
		false,
		false,
		false,
		nil,
	)
	messages, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	for v := range messages {
		str := string(v.Body)
		fmt.Println(str)
	}
}
