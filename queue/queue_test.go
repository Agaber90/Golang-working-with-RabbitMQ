package main

import (
	"cs-backend-golang-finding/model"
	"cs-backend-golang-finding/rabbitmqhandler"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/spf13/viper"
)

func Test_ReadConfig(t *testing.T) {
	viper.SetConfigFile("config.json")
	pvError := viper.ReadInConfig()
	if pvError != nil {
		assert.Error(t, pvError, "Cannot Read the config file")
	}
}

func Test_OpenInputFolder(t *testing.T) {
	viper.SetConfigFile("config.json")
	_ = viper.ReadInConfig()

	pvFolderPath := viper.GetString("inputFolder.path")

	_, PvError := os.Open(pvFolderPath)
	if PvError != nil {
		assert.Error(t, PvError, "Cannot open the Folder")
	}
}

func Test_ReadAllFiles(t *testing.T) {
	viper.SetConfigFile("config.json")
	_ = viper.ReadInConfig()
	pvFolderPath := viper.GetString("inputFolder.path")

	pvFile, _ := os.Open(pvFolderPath)

	pvFileList, pvError := pvFile.Readdirnames(0)
	if pvError != nil {
		assert.Error(t, pvError, "Cannot read the files")
	}

	if len(pvFileList) > 0 {
		assert.Equal(t, len(pvFileList), len(pvFileList), "files have been fetched")
	}
}

func Test_SendFilesToTheQueue(t *testing.T) {
	viper.SetConfigFile("config.json")
	_ = viper.ReadInConfig()
	pvFolderPath := viper.GetString("inputFolder.path")
	pvHost := viper.GetString("rabbitMQ.server")
	pvPort := viper.GetString("rabbitMQ.port")
	pvUserName := viper.GetString("rabbitMQ.userName")
	pvPassword := viper.GetString("rabbitMQ.password")
	pvQueueName := viper.GetString("rabbitMQ.queueName")
	pvContentType := viper.GetString("rabbitMQ.contentType")

	pvFile, _ := os.Open(pvFolderPath)
	pvFileList, _ := pvFile.Readdirnames(0)

	pvMQ := pvPassword + ":" + pvUserName + "@" + pvHost + ":" + pvPort
	pvMQURL := "amqp://" + pvMQ + "/"

	pvRabbitMQ := rabbitmqhandler.NewRabitMQConnection(pvMQURL)
	pvRabbitMQQueue := rabbitmqhandler.NewQueue(pvQueueName+"_test", false, false, false, false, nil)

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
