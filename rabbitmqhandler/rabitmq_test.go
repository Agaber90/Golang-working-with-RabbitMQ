package rabbitmqhandler

import (
	"testing"

	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

func Test_Dial(t *testing.T) {
	pvURL := "amqp://guest:guest@localhost:5672/"
	_, pvError := amqp.Dial(pvURL)
	if pvError != nil {
		assert.Error(t, pvError, "Failed to connect to RabbitMQ")
	}

}

func Test_RMQOpenChannel(t *testing.T) {
	pvURL := "amqp://guest:guest@localhost:5672/"
	pvConn, pvError := amqp.Dial(pvURL)
	if pvError != nil {
		assert.Error(t, pvError, "Failed to connect to RabbitMQ")
	}
	_, pvError = pvConn.Channel()
	if pvError != nil {
		assert.Error(t, pvError, "Failed to open a channel")
	}
}

func Test_DeclareQueue(t *testing.T) {
	pvURL := "amqp://guest:guest@localhost:5672/"
	pvConn, pvError := amqp.Dial(pvURL)
	if pvError != nil {
		assert.Error(t, pvError, "Failed to connect to RabbitMQ")
	}

	pvChannel, pvError := pvConn.Channel()
	if pvError != nil {
		assert.Error(t, pvError, "Failed to open a channel")
	}

	_, pvError = pvChannel.QueueDeclare(
		"test_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if pvError != nil {
		assert.Error(t, pvError, "Failed to declare a queue")
	}
}

func Test_NewRabitMQConnection(t *testing.T) {
	url := "amqp://guest:guest@localhost:5672/"
	pvRMQ := NewRabitMQConnection(url)
	if pvRMQ == nil {
		assert.Nil(t, pvRMQ, "RMQ Object is Nil")
	}
}

func Test_NewQueue(t *testing.T) {
	pvQname := "test_queuw"
	pvDurable := false
	pvAutoDelete := false
	pvExclusive := false
	pvNoWait := false

	pvQueue := NewQueue(pvQname, pvDurable, pvAutoDelete, pvExclusive, pvNoWait, nil)

	if pvQueue == nil {
		assert.Nil(t, pvQueue, "Queueu Object is Nil")
	}
}

func Test_Publish(t *testing.T) {
	pvBody := "testRmq"
	pvExchange := ""
	pvContentType := "text/plain"
	pvMandatory := false
	pvImmediate := false
	pvData := []byte(pvBody)
	url := "amqp://guest:guest@localhost:5672/"
	pvRMQ := NewRabitMQConnection(url)
	if pvRMQ == nil {
		assert.Nil(t, pvRMQ, "RMQ Object is Nil")
	}
	pvQname := "test_queue"
	pvDurable := false
	pvAutoDelete := false
	pvExclusive := false
	pvNoWait := false

	pvQueue := NewQueue(pvQname, pvDurable, pvAutoDelete, pvExclusive, pvNoWait, nil)

	if pvQueue == nil {
		assert.Nil(t, pvQueue, "Queueu Object is Nil")
	}
	pvError := pvRMQ.Publish(
		pvExchange,
		pvContentType,
		pvMandatory,
		pvImmediate,
		pvData,
		pvQueue,
	)

	if pvError != nil {
		assert.Error(t, pvError, "Failed to publish a message")
	}

}

func Test_Consume(t *testing.T) {
	pvConsumer := ""
	pvAutoAck := true
	pvExeclusive := false
	pvNolocal := false
	pvNowait := false

	url := "amqp://guest:guest@localhost:5672/"
	pvRMQ := NewRabitMQConnection(url)
	if pvRMQ == nil {
		assert.Nil(t, pvRMQ, "RMQ Object is Nil")
	}

	pvQname := "test_queue"
	pvDurable := false
	pvAutoDelete := false
	pvExclusive := false
	pvNoWait := false

	pvQueue := NewQueue(pvQname, pvDurable, pvAutoDelete, pvExclusive, pvNoWait, nil)

	if pvQueue == nil {
		assert.Nil(t, pvQueue, "Queueu Object is Nil")
	}
	_, pvError := pvRMQ.ConsumeMsgs(
		pvConsumer,
		pvAutoAck,
		pvExeclusive,
		pvNolocal,
		pvNowait,
		pvQueue,
	)

	if pvError != nil {
		assert.Error(t, pvError, "Failed to register a consumer")
	}
}
