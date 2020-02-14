package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
)

/**
	Author:charlie
	Description:router
	Time:2020-2-11
*/

// JSONParams doc
type JSONParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
var RabbitMQSend = func() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {

		params := &JSONParam{}
		ctx.BindJSON(params)
		//先拿到json数据
		conn,err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
		if err != nil {
			log.Printf("%s,%s","Fail to connect to RabbitMQ",err)
		}
		defer conn.Close()
		channel,err := conn.Channel()
		defer channel.Close()
		if err != nil {
			log.Printf("%s,%s","Fail to open a channel",err)
		}
		defer channel.Close()
		queue,err := channel.QueueDeclare("hell",true,false,false,false,nil,)
		if err != nil {
			log.Printf("%s,%s","Fail to declare queue",err)
		}
		MessageBody := *params
		bytes,err := json.Marshal(params)
		channel.Publish("",queue.Name,false,false,amqp.Publishing{
			DeliveryMode:amqp.Persistent,
			ContentType:"text/plain",
			Body:bytes})
		log.Printf("producer send %s",MessageBody)
		if err!= nil {
			log.Printf("%s,%s","Fail to publish",err)
		}
	})
}
