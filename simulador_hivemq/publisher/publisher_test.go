package main

import (
	"fmt"
	"testing"
	"time"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func TestHiveMQIntegration(t *testing.T) {
	brokerURL := os.Getenv("BROKER_URL")
	username := os.Getenv("BROKER_USERNAME")
	password := os.Getenv("BROKER_PASSWORD")
	topic := "sensor"

	opts := mqtt.NewClientOptions().
		AddBroker(brokerURL).
		SetClientID("integration-test-client").
		SetUsername(username).
		SetPassword(password)

	// Criando um cliente MQTT para a publicação da mensagem
	publisher := mqtt.NewClient(opts)
	if token := publisher.Connect(); token.Wait() && token.Error() != nil {
		t.Fatalf("Publisher couldn't connect to HiveMQ: %v", token.Error())
	}
	defer publisher.Disconnect(250)

	// Criando um cliente MQTT para a subscrição
	subscriber := mqtt.NewClient(opts)
	if token := subscriber.Connect(); token.Wait() && token.Error() != nil {
		t.Fatalf("Subscriber couldn't connect to HiveMQ: %v", token.Error())
	}
	defer subscriber.Disconnect(250)

	received := make(chan struct{})
	payload := "Hello HiveMQ"

	// Inscrevendo-se para receber a mensagem
	if token := subscriber.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		if string(msg.Payload()) != payload {
			t.Errorf("Expected %s but got %s", payload, msg.Payload())
		} else {
			fmt.Println("Message received successfully")
		}
		close(received)
	}); token.Wait() && token.Error() != nil {
		t.Fatalf("Subscriber couldn't subscribe to topic: %v", token.Error())
	}

	// Publicando a mensagem
	if token := publisher.Publish(topic, 0, false, payload); token.Wait() && token.Error() != nil {
		t.Fatalf("Publisher couldn't publish to topic: %v", token.Error())
	}

	// Esperando a mensagem ser recebida ou timeout
	select {
	case <-received:
		// Mensagem recebida
	case <-time.After(5 * time.Second):
		t.Fatal("Timeout waiting for the message")
	}
}
