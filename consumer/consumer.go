package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer Application")

	//Dial accepts a string in the AMQP URI format and returns a new Connection over TCP using PlainAuth
	conn , err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err!=nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()

	channel , err := conn.Channel()
	if err!=nil {
		fmt.Println(err)
		panic(err)
	}
	defer channel.Close()

	msgs, err := channel.Consume("testQueue", "", true, false,false,false,nil)
	//returns a channel
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved messages: %s\n",d.Body)
		}
	}()
	fmt.Println("Succesfully connect to rabbitmq server instance")
	fmt.Println("Waiting for messages.....")
	<-forever
}