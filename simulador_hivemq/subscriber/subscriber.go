package main 

import (
	"fmt"
	"os"
	"time"
	
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Recebido: %s do tópico: %s\n", msg.Payload(), msg.Topic())
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Conexão perdida: %v\n", err)
}

func main() {
	// Configuração do cliente MQTT
	opts := mqtt.NewClientOptions().
		AddBroker("tcp://734ec015f104410abedf1ab91c9ef915.s1.eu.hivemq.cloud:1883").
		SetClientID("").
		SetDefaultPublishHandler(messagePubHandler).
		SetConnectionLostHandler(connectLostHandler)

	// Criação do cliente MQTT com as opções configuradas
	client := mqtt.NewClient(opts)

	// Conectar ao broker
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Erro ao conectar:", token.Error())
		os.Exit(1)
	}

	// Inscrever-se no tópico "sensor/data"
	if token := client.Subscribe("sensor/data", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println("Erro ao se inscrever:", token.Error())
		os.Exit(1)
	}

	fmt.Println("Subscriber está rodando, pressione CTRL+C para sair...")
	
	// Mantém o subscriber rodando
	for {
		time.Sleep(1 * time.Second)
	}
}