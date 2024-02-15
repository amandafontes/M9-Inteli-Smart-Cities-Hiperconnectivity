import paho.mqtt.client as mqtt
import random
import json
import time

# Configuração do cliente
client = mqtt.Client(mqtt.CallbackAPIVersion.VERSION1, client_id="")

# Conecte ao broker
client.connect("localhost", 1891, 60)

# Leitura do arquivo JSON
with open("sensor.json", "r") as f:
    data = json.load(f)

gases_detectaveis = data["gases_detectaveis"]

# Loop para publicar mensagens continuamente de maneira aleatória
try:
    while True:
        index = random.randint(0, len(gases_detectaveis) - 1)
        
        message = json.dumps(gases_detectaveis[index])
        client.publish("sensor/data", message)
        print(f"Publicado: {message}")
        time.sleep(2)
except KeyboardInterrupt:
    print("Publicação encerrada")
    client.disconnect()