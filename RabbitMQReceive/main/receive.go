package main

import (
	"github.com/streadway/amqp"
	"log"
)
/**
	Author:charlie
	Description:consumer receive message to queue
	Time:2020-2-10
*/
func main()  {
	conn,err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Printf("%s,%s","fail to connect RabbitMQ",err)
	}
	defer conn.Close()
	channel,err := conn.Channel()
	if err != nil {
		log.Printf("%s,%s","fail to open a channel",err)
	}
	defer channel.Close()
	queue,err := channel.QueueDeclare("hell",true,false,false,false,nil)
	if err != nil {
		log.Printf("%s,%s","fail to declare a queue",err)
	}
	msg,err := channel.Consume(queue.Name,"",false,false,false,false,nil)
	if err != nil {
		log.Printf("%s,%s","fail to regigter",err)
	}
	listen := make(chan bool)
	go func() {
		for d := range msg {
			log.Printf("Received message is %s",d.Body)
			d.Ack(false)
		}
	}()
	log.Println("waiting for message")
	<-listen
}
