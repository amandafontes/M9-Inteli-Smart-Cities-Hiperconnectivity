package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
	"crypto/tls"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func PublishData(client mqtt.Client, topic string, payload interface{}) {
	message, _ := json.Marshal(payload)
	token := client.Publish(topic, 1, false, message)
	token.Wait()

	fmt.Printf("Publicado: %s\n", message)
}

func main() {

	brokerURL := os.Getenv("BROKER_URL")
	username := os.Getenv("BROKER_USERNAME")
	password := os.Getenv("BROKER_PASSWORD")

	// Configuração do cliente MQTT
	opts := mqtt.NewClientOptions().AddBroker(brokerURL).SetClientID("unique-client-id")
	opts.SetUsername(username)
	opts.SetPassword(password)

	// Configuração do TLS
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	opts.SetTLSConfig(tlsConfig)

	client := mqtt.NewClient(opts)

	// Conectar ao broker
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Leitura do arquivo JSON
	file, err := os.ReadFile("../sensor.json")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		os.Exit(1)
	}

	var data map[string][]map[string]interface{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Erro ao fazer o parse do JSON:", err)
		os.Exit(1)
	}

	gasesDetectaveis := data["gases_detectaveis"]
	rand.Seed(time.Now().UnixNano())

	// Loop para publicar mensagens continuamente de maneira aleatória
	for {
		index := rand.Intn(len(gasesDetectaveis))
		message, _ := json.Marshal(gasesDetectaveis[index])
		token := client.Publish("sensor", 1, false, message)
		token.Wait()

		fmt.Printf("Publicado: %s\n", message)
		time.Sleep(time.Second)
	}
}
