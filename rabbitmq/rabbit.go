package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type MQ interface {
	Bind(exchange string)
	Send(queue string, body interface{})
	Publish(exchange string, body interface{})
	Consume() <-chan amqp.Delivery
	Close()
}

type RabbitMQ struct {
	Name     string
	channel  *amqp.Channel
	exchange string
}

func New(addr string) MQ {
	conn, e := amqp.Dial(addr)
	if e != nil {
		log.Fatal(e)
	}

	ch, e := conn.Channel()
	if e != nil {
		log.Fatal(e)
	}

	q, e := ch.QueueDeclare(
		"",    // name
		false, // durable
		true,  // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if e != nil {
		log.Fatal(e)
	}

	mq := new(RabbitMQ)
	mq.channel = ch
	mq.Name = q.Name

	return mq
}

func (r *RabbitMQ) Bind(exchange string) {
	if e := r.channel.QueueBind(
		r.Name, // queue name
		"",     // routing key
		exchange,
		false, nil); e != nil {
		log.Fatal(e)
	}

	r.exchange = exchange
}

func (r *RabbitMQ) Send(queue string, body interface{}) {
	str, e := json.Marshal(body)
	if e != nil {
		log.Fatal(e)
	}

	if e := r.channel.Publish("", queue, false, false,
		amqp.Publishing{ReplyTo: r.Name, Body: []byte(str)}); e != nil {
		log.Fatal(e)
	}
}

func (r *RabbitMQ) Publish(exchange string, body interface{}) {
	str, e := json.Marshal(body)
	if e != nil {
		log.Fatal(e)
	}

	if e := r.channel.Publish(
		exchange, "", false, false, amqp.Publishing{
			ReplyTo: r.Name,
			Body:    []byte(str),
		}); e != nil {
		log.Fatal(e)
	}
}

func (r *RabbitMQ) Consume() <-chan amqp.Delivery {
	c, e := r.channel.Consume(r.Name, "", true, false, false, false, nil)
	if e != nil {
		log.Fatal(e)
	}

	return c
}

func (r *RabbitMQ) Close() {
	r.channel.Close()
}
