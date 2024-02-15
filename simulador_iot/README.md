<h2>Semana 1 | Criação de Simulador IoT</h2>

O presente diretório é destinado à entrega da atividade referente à criação de um simulador de dispositivos IoT.

<h3>Introdução à atividade</h3>

<p>O objetivo é criar um simulador de dispositivos IoT utilizando o protocolo MQTT através do uso da biblioteca Eclipse Paho. </p>

<h3>Implementação</h3>

Para a implementação da atividade, foram desenvolvidos um script para publicar os dados do sensor, <code>publisher.py</code>, e um script em que é realizada a subscrição ao tópico no qual as mensagens são publicadas, <code>subscriber.py</code>. Além disso, o arquivo <code>mosquitto.conf</code> possibilita estabelecer a conexão com o broket, a fim de que o protocolo MQTT seja concretizado. O publicador lê o conteúdo contido em <code>sensor.json</code> e publica, de forma aleatória, as informações simuladas em um tópico nomeado <code>sensor/data</code>.

<h3>Demonstração</h3>