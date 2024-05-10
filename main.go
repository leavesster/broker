package main

import (
	"fmt"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

func main() {

	brokerPort := "47688"
	if len(os.Args) > 1 {
		brokerPort = os.Args[1]
	}

	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:" + brokerPort)
	opts.SetClientID("go_mqtt_client")
	opts.SetDefaultPublishHandler(messagePubHandler)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	subscribeTopic(client, "#")

	for {
		time.Sleep(1 * time.Second)
	}
}

func subscribeTopic(client MQTT.Client, topic string) {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s\n", topic)
}
