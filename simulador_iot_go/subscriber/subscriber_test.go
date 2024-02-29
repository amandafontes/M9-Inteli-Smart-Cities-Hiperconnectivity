package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type MockMQTTClient struct {
	mock.Mock
}

func (m *MockMQTTClient) Connect() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockMQTTClient) Subscribe(topic string, qos byte, callback MQTT.MessageHandler) error {
	args := m.Called(topic, qos, callback)
	return args.Error(0)
}

func (m *MockMQTTClient) Disconnect(quiesce uint) error {
	args := m.Called(quiesce)
	return args.Error(0)
}

func TestSubscriberRun(t *testing.T) {
	// Configurar o mock do cliente MQTT
	mockClient := new(MockMQTTClient)
	mockClient.On("Connect").Return(nil)
	mockClient.On("Subscribe", "sensor/data", byte(1), mock.Anything).Return(nil)
	mockClient.On("Disconnect", uint(250)).Return(nil)

	// Criar uma instância do Subscriber com o mock do cliente MQTT
	subscriber := NewSubscriber(mockClient)

	// Executar o método Run do Subscriber
	go subscriber.Run()

	// Aguardar por um tempo para simular a execução
	time.Sleep(2 * time.Second)

	// Verificar se os métodos do mock foram chamados
	mockClient.AssertExpectations(t)
}
