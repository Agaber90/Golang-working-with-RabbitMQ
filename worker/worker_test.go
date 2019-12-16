package main

import (
	"cs-backend-golang-finding/categoryhandler"
	"cs-backend-golang-finding/model"
	"cs-backend-golang-finding/rabbitmqhandler"
	"encoding/json"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func Test_ReadConfig(t *testing.T) {
	viper.SetConfigFile("config.json")
	pvError := viper.ReadInConfig()
	if pvError != nil {
		assert.Error(t, pvError, "Cannot Read the config file")
	}
}

func Test_OpenOutputFolderExist(t *testing.T) {
	viper.SetConfigFile("config.json")
	_ = viper.ReadInConfig()

	pvFolderPath := viper.GetString("outputFolder.path")

	_, PvError := os.Open(pvFolderPath)
	if PvError != nil {
		assert.Error(t, PvError, "Cannot open the Folder")
	}
}

func Test_OpenOutputFolderIsNotExist(t *testing.T) {
	viper.SetConfigFile("config.json")
	_ = viper.ReadInConfig()

	pvFolderPath := viper.GetString("outputFolder.path1")

	_, PvError := os.Open(pvFolderPath)
	if PvError != nil {
		assert.Error(t, PvError, "Cannot open the Folder")
	}
}

func Test_AddMessagesToOutPutFolder(t *testing.T) {
	viper.SetConfigFile("config.json")
	_ = viper.ReadInConfig()
	pvHost := viper.GetString("rabbitMQ.server")
	pvPort := viper.GetString("rabbitMQ.port")
	pvUserName := viper.GetString("rabbitMQ.userName")
	pvPassword := viper.GetString("rabbitMQ.password")
	pvQueueName := viper.GetString("rabbitMQ.queueName")
	pvOutputFolder := viper.GetString("outputFolder.path")

	pvFindingBody := model.FindingBody{}
	pvCategory := model.CategoriesBody{}

	pvMQ := pvPassword + ":" + pvUserName + "@" + pvHost + ":" + pvPort
	pvMQURL := "amqp://" + pvMQ + "/"
	pvRabbitMQ := rabbitmqhandler.NewRabitMQConnection(pvMQURL)

	pvRabbitMQQueue := rabbitmqhandler.NewQueue(
		pvQueueName+"_test",
		false,
		false,
		false,
		false,
		nil,
	)

	pvMessages, _ := pvRabbitMQ.ConsumeMsgs(
		"",
		true,
		false,
		false,
		false,
		pvRabbitMQQueue,
	)

	pvForever := make(chan bool)

	go func() {
		for msg := range pvMessages {
			_ = json.Unmarshal([]byte(msg.Body), &pvFindingBody)
			categoryhandler.HandleCategory(&pvFindingBody, &pvCategory, pvOutputFolder)

		}
	}()
	defer pvRabbitMQ.Conn.Close()
	defer pvRabbitMQ.Channel.Close()
	<-pvForever
}
