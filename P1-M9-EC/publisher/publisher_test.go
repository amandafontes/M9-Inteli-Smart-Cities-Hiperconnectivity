package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func TestMQTTConnection(t *testing.T) {
	broker := "tcp://localhost:1891"
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID("testConnection")
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Fatalf("Erro ao conectar ao broker: %v", token.Error())
	}
	defer client.Disconnect(250)
	fmt.Println("Conexão bem-sucedida com o broker")
}

func TestMQTTSubscription(t *testing.T) {
	broker := "tcp://localhost:1891"
	topic := "temperature"
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID("subscription")
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Fatalf("Erro ao conectar ao broker: %v", token.Error())
	}
	defer client.Disconnect(250)

	if token := client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		t.Fatalf("Erro ao subscrever no tópico %s: %v", topic, token.Error())
	}
	fmt.Println("Subscrição bem-sucedida ao tópico")
}

func TestMQTTMessageReception(t *testing.T) {
	topic := "temperature"
	broker := "tcp://localhost:1891"
	var messageCount int
	testDuration := 10 * time.Second
	expectedMinimumMessages := 5

	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID("testMessageReception")
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Fatalf("Erro ao conectar ao broker: %v", token.Error())
	}

	if token := client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		messageCount++
		fmt.Printf("Mensagem recebida: %s\n", msg.Payload())

		var receivedData map[string]interface{}
		if err := json.Unmarshal(msg.Payload(), &receivedData); err != nil {
			t.Errorf("Erro ao fazer parse da mensagem JSON recebida: %v", err)
		}
	}); token.Wait() && token.Error() != nil {
		t.Fatalf("Erro ao subscrever no tópico %s: %v", topic, token.Error())
	}

	time.Sleep(testDuration)

	if messageCount < expectedMinimumMessages {
		t.Errorf("Recebidas menos mensagens (%d) do que o esperado (%d)", messageCount, expectedMinimumMessages)
	} else {
		fmt.Printf("Total de mensagens recebidas: %d em %s. Taxa: %.2f msgs/segundo\n", messageCount, testDuration, float64(messageCount)/testDuration.Seconds())
	}
}
