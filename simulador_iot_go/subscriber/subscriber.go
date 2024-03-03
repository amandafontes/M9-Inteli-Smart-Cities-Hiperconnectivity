package main

import (
	"fmt"
	"os"
	"os/signal"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type MQTTClient interface {
	Connect() error
	Subscribe(topic string, qos byte, callback MQTT.MessageHandler) error
	Disconnect(quiesce uint) error
}

type Subscriber struct {
	Client MQTTClient
}

var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Recebido: %s do tópico: %s\n", msg.Payload(), msg.Topic())
}

func NewSubscriber(client MQTTClient) *Subscriber {
	return &Subscriber{
		Client: client,
	}
}

func (s *Subscriber) Run() {
	// O restante do código do Run permanece o mesmo...
}

func main() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_subscriber")
	opts.SetDefaultPublishHandler(messagePubHandler)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("sensor/data", 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	fmt.Println("Subscriber está rodando. Pressione CTRL+C para sair.")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	client.Disconnect(250)
	fmt.Println("Subscriber desconectado.")
}
