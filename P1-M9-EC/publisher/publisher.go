package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func PublishData(client mqtt.Client, topic string, payload interface{}) {
	message, _ := json.Marshal(payload)
	token := client.Publish(topic, 1, false, message)
	token.Wait()

	fmt.Printf("Publicado: %s\n", message)
}

func main() {

	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("subscription")
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	file, err := os.ReadFile("../data.json")
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

	temperatureData := data["dados"]

	for {
		index := rand.Intn(len(temperatureData))
		message, _ := json.Marshal(temperatureData[index])
		token := client.Publish("temperature", 1, false, message)
		token.Wait()

		fmt.Printf("Publicado: %s\n", message)
		time.Sleep(2 * time.Second)
	}
}