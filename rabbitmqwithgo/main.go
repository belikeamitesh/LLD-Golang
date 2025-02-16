package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {

	//amqp - advanced message queue protocol
	fmt.Println("Welcome to rabbitmq connection")

	//Dial accepts a string in the AMQP URI format and returns a new Connection over TCP using PlainAuth
	conn , err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err!=nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()
	fmt.Println("Successfully connect to rabbitmq instance")

	//Channel opens a unique, concurrent server channel to process the bulk of AMQP messages
	channel , err := conn.Channel()
	if err!=nil {
		fmt.Println(err)
		panic(err)
	}
	defer channel.Close()

	//declares a queue to hold messages and deliver to consumers
	//func (ch *amqp.Channel) QueueDeclare(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args amqp.Table) (amqp.Queue, error)
	queue,err := channel.QueueDeclare("testQueue",false, false, false, false, nil)
	if err!=nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(queue)

	//Publish(exchange string, key string, mandatory bool, immediate bool, msg amqp.Publishing) error
	err = channel.Publish("","testQueue",false,false,amqp.Publishing{
		ContentType: "text/plain",
		Body: []byte("Hello world"),
	})

	if err!=nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Succesfully Published message to queue")
}