package rabbitmqhandler

import (
	"log"

	"github.com/streadway/amqp"
)

//RabbitMQ that holds the url and connection
type RabbitMQ struct {
	url     string
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

//NewRabitMQConnection to intialize a connection
func NewRabitMQConnection(URL string) *RabbitMQ {
	pvRMQ := new(RabbitMQ)
	pvRMQ.url = URL
	pvConn, pvError := amqp.Dial(pvRMQ.url)
	if pvError != nil {
		log.Fatal(pvError)
		panic(pvError)
	}
	pvRMQ.Conn = pvConn

	return pvRMQ
}

//Queue that holds the queue data
type Queue struct {
	queueName   string
	durable     bool
	autotDelete bool
	exclusive   bool
	noWait      bool
	args        amqp.Table
}

//NewQueue to Intialize a new queue
func NewQueue(parQname string, parDurable, ParAutoDelete, parExclusive, parNoWait bool, ParArgs amqp.Table) *Queue {
	pvQueue := new(Queue)
	pvQueue.queueName = parQname
	pvQueue.durable = parDurable
	pvQueue.autotDelete = ParAutoDelete
	pvQueue.exclusive = parExclusive
	pvQueue.noWait = parNoWait
	pvQueue.args = ParArgs
	return pvQueue
}

//Publish to publish a message to RMQ
func (r *RabbitMQ) Publish(parExchange, parContentType string, parMandatory, parImmediate bool, parBody []byte, parQuque *Queue) error {

	pvChannel, pvError := r.Conn.Channel()
	if pvError != nil {
		log.Fatal(pvError)
		panic(pvError)
	}

	r.Channel = pvChannel

	pvQueue, pvError := pvChannel.QueueDeclare(parQuque.queueName, parQuque.durable, parQuque.autotDelete, parQuque.exclusive, parQuque.noWait, parQuque.args)
	if pvError != nil {
		log.Fatal(pvError)
		panic(pvError)
	}

	pvError = pvChannel.Publish(
		parExchange,
		pvQueue.Name,
		parMandatory,
		parImmediate,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  parContentType,
			Body:         parBody,
		})

	return pvError
}

//ConsumeMsgs to Consume messages from RMQ
func (r *RabbitMQ) ConsumeMsgs(parConsumer string, parAutoAck, parExeclusive, parNolocal, parNowait bool, parQuque *Queue) (<-chan amqp.Delivery, error) {
	pvChannel, pvError := r.Conn.Channel()
	if pvError != nil {
		log.Fatal(pvError)
		panic(pvError)
	}

	r.Channel = pvChannel

	pvQueue, pvError := pvChannel.QueueDeclare(parQuque.queueName, parQuque.durable, parQuque.autotDelete, parQuque.exclusive, parQuque.noWait, parQuque.args)
	if pvError != nil {
		log.Fatal(pvError)
		panic(pvError)
	}

	if pvError != nil {
		log.Fatal(pvError)
		panic(pvError)
	}

	pvMessages, pvError := pvChannel.Consume(
		pvQueue.Name,
		parConsumer,
		parAutoAck,
		parExeclusive,
		parNolocal,
		parNowait,
		parQuque.args,
	)

	if pvError != nil {
		log.Fatal(pvError)
		panic(pvError)
	}
	return pvMessages, pvError
}
