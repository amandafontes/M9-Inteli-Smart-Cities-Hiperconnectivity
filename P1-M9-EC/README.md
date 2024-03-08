<h2>Atividade Prática | P1-M9-EC</h2>

<p>Diretório destinado à entrega da primeira prova do módulo 9 do curso de Engenharia da Computação.</p>

<h3>Introdução</h3>

O objetivo da atividade prática é o desenvolvimento e implementação de um sistema de monitoramento simulado de temperatura em uma cadeia de supermercados.

<i>Para simular a operação do sistema, você deve criar um pequeno simulador (publisher) capaz de enviar mensagens simulando os sensores de temperatura. </i>

<h3>Implementação</h3>

Para o desenvolvimento da atividade, foi utilizada a linguagem <code>Go</code>, a biblioteca <code>Paho MQTT</code> e o broker local <code>mosquitto</code>.

Para a implementação da atividade, foi desenvolvido um script <code>publisher.go</code>, que atua como publicador em tempo real dos dados de temperatura que desejamos monitorar.

O script, primeiramente, instancia o cliente MQTT e estabelece conexão com o broker local. Posteriormente, ele lê o arquivo em que estão contidos os dados simulados que desejamos publicar. Por fim, de forma randômica e periódica, realiza a publicação dos dados simulados de temperatura.

O script de testes, <code>publisher_test.go</code>, garante, atualmente, os seguintes aspectos:

**1.** A conexão bem-sucedida do cliente MQTT com o broker local, verificada por meio da função <code>TestMQTTSubscription</code>.

**2.** A subscrição bem-sucedida ao tópico por meio do qual o publisher envia a mensagem, bem como o recebimento da mensagem enviada, verificado por meio da função <code>TestMQTTSubscription</code>.

**3.** Se a periodicidade com a qual as mensagens são enviadas pelo publisher, verificada por meio da função <code>TestMQTTMessageReception</code>, está de acordo com o esperado.

<h3>Execução</h3>

Para executar o sistema de forma bem-sucedida, é necessário abrir um terminal no presente diretório, executar o broker, executar o publicador e, por fim, executar o script de testes. Para a concretização desse objetivo, recomenda-se o uso de três terminais paralelos:

**1.** Execução do broker local
```bash
mosquitto -c mosquitto.conf
```

**2.** Execução do publisher
```bash
go mod tidy
go run publisher.go
```

**3.** Execução do script de testes
```bash
go test
```

<h3>Próximos passos</h3>

Ao desenvolver a atividade, tive algumas dificuldades no que concerne a determinados requisitos evidenciados no enunciado da prova. Abaixo, encontram-se descritos os elementos pendentes no momento.

- O desenvolvimento de um sistema de alertas capaz de notificar o usuário quando a temperatura identificada for atípica.

- A implementação de testes relacionados ao sistema de alarmes e à verificação do Quality of Service de nível 1.

- A formatação esperada para o payload na saída exibida pelo terminal.