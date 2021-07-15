package main

import (
	"fmt"
	"github.com/djumanoff/amqp"
)

func main() {
	config := amqp.Config{
		Host: "localhost",
		Port: 5672,
	}
	sess := amqp.NewSession(config)
	err := sess.Connect()
	if err != nil {
		panic(err)
		return
	}
	clt, err := sess.Client(amqp.ClientConfig{})
	if err != nil {
		panic(err)
		return
	}
	res, err := clt.Call("lesson18.GetNumber", amqp.Message{})
	fmt.Println(string(res.Body))
}
