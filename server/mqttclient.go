package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mqttClient mqtt.Client
var mqttConnected bool = false

func mqttConnect(mqttBroker string, username string, password string) error {
	var messageSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		println("Received message on topic:", msg.Topic(), "with payload:\n", string(msg.Payload()))
	}

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

func mqttSubscribe(topic string, callback mqtt.MessageHandler) error {
	if token := mqttClient.Subscribe(topic, 0, callback); token.Wait() && token.Error() != nil {
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
