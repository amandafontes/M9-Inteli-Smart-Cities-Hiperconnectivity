package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Test(t *testing.T) {
	topic := "sensor/data"
	broker := "tcp://localhost:1891"
	var messageCount int
	testDuration := 10 * time.Second // Ajuste conforme a taxa de disparo esperada do seu simulador
	expectedMinimumMessages := 5     // Ajuste de acordo com a taxa de disparo esperada e a duração do teste

	// Configuração do cliente MQTT para o subscriber de teste
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID("testSubscriber")
	client := mqtt.NewClient(opts)

	// Conectar ao broker
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Fatalf("Erro ao conectar ao broker: %v", token.Error())
	}
	fmt.Printf("Conexão bem-sucedida com o broker\n")

	defer client.Disconnect(250)

	// Subscrever ao tópico e contar mensagens recebidas
	if token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		messageCount++
		fmt.Printf("Mensagem recebida: %s\n", msg.Payload())

		// Aqui, você pode adicionar verificações adicionais sobre o payload, se necessário
		var receivedData map[string]interface{}
		if err := json.Unmarshal(msg.Payload(), &receivedData); err != nil {
			t.Errorf("Erro ao fazer parse da mensagem JSON recebida: %v", err)
		}
	}); token.Wait() && token.Error() != nil {
		t.Fatalf("Erro ao subscrever no tópico %s: %v", topic, token.Error())
	}

	// Aguardar pelo período de teste
	time.Sleep(testDuration)

	// Verificar se o número de mensagens recebidas está dentro da expectativa
	if messageCount < expectedMinimumMessages {
		t.Errorf("Recebidas menos mensagens (%d) do que o esperado (%d)", messageCount, expectedMinimumMessages)
	}

	// Calcula a taxa de mensagens recebidas por segundo para informação
	fmt.Printf("Total de mensagens recebidas: %d em %s. Taxa: %.2f msgs/segundo\n", messageCount, testDuration, float64(messageCount)/testDuration.Seconds())
}
