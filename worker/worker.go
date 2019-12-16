package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"cs-backend-golang-finding/categoryhandler"
	"cs-backend-golang-finding/model"
	"cs-backend-golang-finding/rabbitmqhandler"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")
	pvError := viper.ReadInConfig()
	if pvError != nil {
		panic(pvError)
	}

	if viper.GetBool("debug") {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {

	pvHost := viper.GetString("rabbitMQ.server")
	pvPort := viper.GetString("rabbitMQ.port")
	pvUserName := viper.GetString("rabbitMQ.userName")
	pvPassword := viper.GetString("rabbitMQ.password")
	pvQueueName := viper.GetString("rabbitMQ.queueName")
	pvOutputFolder := viper.GetString("outputFolder.path")

	/*This responsable to consume messages at the same time based on the configurable worker
	*This configurable I added 5 as the document mentiond to consume 5 messages at the same time
	 */
	concurentWorker := viper.GetInt("rabbitMQ.worker")

	pvFindingBody := model.FindingBody{}
	pvCategory := model.CategoriesBody{}

	pvMQ := pvPassword + ":" + pvUserName + "@" + pvHost + ":" + pvPort
	pvMQURL := "amqp://" + pvMQ + "/"
	pvRabbitMQ := rabbitmqhandler.NewRabitMQConnection(pvMQURL)

	pvRabbitMQQueue := rabbitmqhandler.NewQueue(
		pvQueueName,
		false,
		false,
		false,
		false,
		nil,
	)

	pvMessages, pvError := pvRabbitMQ.ConsumeMsgs(
		"",
		false,
		false,
		false,
		false,
		pvRabbitMQQueue,
	)
	if pvError != nil {
		log.Fatal(pvError)
		panic(pvError)
	}
	pvForever := make(chan bool)
	var wg sync.WaitGroup
	for i := 0; i <= concurentWorker; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case msg := <-pvMessages:
					_ = json.Unmarshal([]byte(msg.Body), &pvFindingBody)
					categoryhandler.HandleCategory(&pvFindingBody, &pvCategory, pvOutputFolder)
					msg.Ack(false)
				}
			}
		}()
	}
	wg.Wait()
	<-pvForever
}
