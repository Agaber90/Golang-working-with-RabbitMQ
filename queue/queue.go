package main

import (
	"cs-backend-golang-finding/model"
	"cs-backend-golang-finding/rabbitmqhandler"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

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

	pvFolderPath := viper.GetString("inputFolder.path")
	pvHost := viper.GetString("rabbitMQ.server")
	pvPort := viper.GetString("rabbitMQ.port")
	pvUserName := viper.GetString("rabbitMQ.userName")
	pvPassword := viper.GetString("rabbitMQ.password")
	pvQueueName := viper.GetString("rabbitMQ.queueName")
	pvContentType := viper.GetString("rabbitMQ.contentType")

	pvFile, PvError := os.Open(pvFolderPath)
	if PvError != nil {
		log.Fatalf("failed opening directory: %s", PvError)
		panic(PvError)
	}

	pvFileList, pvError := pvFile.Readdirnames(0)
	if pvError != nil {
		log.Fatalf("failed list the files: %s", PvError)
		panic(pvError)
	}

	pvMQ := pvPassword + ":" + pvUserName + "@" + pvHost + ":" + pvPort
	pvMQURL := "amqp://" + pvMQ + "/"

	pvRabbitMQ := rabbitmqhandler.NewRabitMQConnection(pvMQURL)
	pvRabbitMQQueue := rabbitmqhandler.NewQueue(
		pvQueueName,
		false,
		false,
		false,
		false,
		nil)

	for _, pvFielName := range pvFileList {
		pvJSONFile, _ := ioutil.ReadFile(pvFolderPath + pvFielName)
		pvFinding := model.FindingBody{}
		_ = json.Unmarshal([]byte(pvJSONFile), &pvFinding)
		pvBodyResponse, _ := json.Marshal(pvFinding)
		pvError := pvRabbitMQ.Publish(
			"",
			pvContentType,
			false,
			false,
			pvBodyResponse,
			pvRabbitMQQueue,
		)
		if pvError != nil {
			fmt.Print("Error: ", pvError)
		}
	}
	defer pvRabbitMQ.Channel.Close()
	defer pvRabbitMQ.Conn.Close()
}
