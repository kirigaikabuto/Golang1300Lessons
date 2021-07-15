package main

import (
	"github.com/streadway/amqp"
)

func main() {
	//producer
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
	if err != nil {
		panic(err)
		return
	}
	message := "Hello my name is 132131112313"
	err = channel.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		panic(err)
		return
	}

}
