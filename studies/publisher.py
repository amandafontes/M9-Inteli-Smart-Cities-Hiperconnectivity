import paho.mqtt.client as mqtt
import time

# Configuração do cliente
client = mqtt.Client(mqtt.CallbackAPIVersion.VERSION1, client_id="")

# Conecte ao broker
client.connect("localhost", 1891, 60)

# Loop para publicar mensagens continuamente
try:
    while True:
        message = "Hello MQTT " + time.strftime("%H:%M:%S")
        client.publish("test/topic", message)
        print(f"Publicado: {message}")
        time.sleep(2)
except KeyboardInterrupt:
    print("Publicação encerrada")

client.disconnect()