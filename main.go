package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const configFile = "./backend_config.json"

var configuration *ConfigFile
var wxServer *MqServer
var timeSec time.Duration

func init() {
	file, err := os.Open(configFile)
	decoder := json.NewDecoder(file)
	configuration = &ConfigFile{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(configuration)
	wxServer = NewRedisMqServer(configuration.AppId, configuration.AppSecret, configuration.MqAddress)
	timeSec = time.Duration(configuration.delay) * time.Second
}

func main() {
	for {
		val := wxServer.Mq.Poll("key", timeSec)
		wxServer.HandleMessage(val)
	}
}
