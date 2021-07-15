package utils

import "github.com/djumanoff/amqp"

func AMQPGetNumber() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		res := GetNumber()
		response := &amqp.Message{Body: []byte(res)}
		return response
	}
}
