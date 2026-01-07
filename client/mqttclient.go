package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mqttClient mqtt.Client
var mqttConnected bool = false

// var mqttBroker string = "tcp://45.117.179.134:1883"

func mqttConnect(mqttBroker string, username string, password string) error {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(mqttBroker)
	opts.SetClientID("aes-client")
	opts.SetUsername(username)
	opts.SetPassword(password)
	mqttClient = mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	mqttConnected = true
	return nil
}

func mqttPublish(topic string, payload []byte) error {
	if token := mqttClient.Publish(topic, 0, false, payload); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
func mqttDisconnect() {
	if mqttConnected {
		mqttClient.Disconnect(250)
		mqttConnected = false
	}
}
