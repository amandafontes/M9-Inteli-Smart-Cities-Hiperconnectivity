package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func TestPublishMessage(t *testing.T) {
	topic := "sensor/data"
	broker := "tcp://localhost:1891"

	// Configuração do cliente MQTT para o subscriber de teste
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID("testSubscriber")
	client := mqtt.NewClient(opts)

	// Conectar ao broker
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Fatalf("Erro ao conectar ao broker: %v", token.Error())
	}

	defer client.Disconnect(250)

	// Canal para receber mensagens
	msgChan := make(chan mqtt.Message)

	// Subscrever ao tópico
	if token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		msgChan <- msg
	}); token.Wait() && token.Error() != nil {
		t.Fatalf("Erro ao subscrever no tópico %s: %v", topic, token.Error())
	}

	// Aguardar por uma mensagem (com timeout)
	select {
	case msg := <-msgChan:
		fmt.Printf("Mensagem recebida: %s\n", msg.Payload())

		var receivedData map[string]interface{}
		if err := json.Unmarshal(msg.Payload(), &receivedData); err != nil {
			t.Fatalf("Erro ao fazer parse da mensagem JSON recebida: %v", err)
		}
		// Adicione verificações específicas aqui baseadas em seu schema JSON
	case <-time.After(5 * time.Second): // Ajuste este timeout conforme necessário
		t.Fatal("Timeout: nenhuma mensagem recebida")
	}
}
