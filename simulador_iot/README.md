<h2>Semana 1 | Criação de Simulador IoT</h2>

O presente diretório é destinado à entrega da atividade referente à criação de um simulador de dispositivos IoT.

<h3>Introdução à atividade</h3>

<p>O objetivo é criar um simulador de dispositivos IoT utilizando o protocolo MQTT através do uso da biblioteca Eclipse Paho.</p>

<h3>Implementação</h3>

Para a implementação da atividade, foram desenvolvidos um script para publicar os dados do sensor, <code>publisher.py</code>, e um script em que é realizada a subscrição ao tópico no qual as mensagens são publicadas, <code>subscriber.py</code>. Além disso, o arquivo <code>mosquitto.conf</code> possibilita estabelecer a conexão com o broket, a fim de que o uso do protocolo MQTT seja concretizado. O publicador lê o conteúdo contido em <code>sensor.json</code> e publica, de forma aleatória, as informações simuladas em um tópico nomeado <code>sensor/data</code>.

São necessários três terminais para executar o sistema. O primeiro será responsável pela inicialização do Mosquitto, o Broker MQTT local escolhido para a atividade.

```shell
mosquitto -c mosquitto.conf
```

O segundo terminal será responsável pela execução do publisher.

```shell
python3 publisher.py
```

O terceiro terminal, por fim, será dedicado à execução do subscriber.

```shell
python3 subscriber.py
```

<h3>Demonstração</h3>

Abaixo, é possível visualizar o vídeo demonstrativo gravado para a atividade.

https://github.com/amandafontes/M9-Inteli-Smart-Cities-Hiperconnectivity/blob/main/simulador_iot/A1-M9-EC.mp4