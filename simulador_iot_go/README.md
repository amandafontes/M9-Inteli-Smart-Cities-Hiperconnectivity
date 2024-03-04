<h2>Semana 2 | Testes para Simulação de Dispositivos IoT</h2>

O presente diretório é destinado à entrega da atividade referente à criação de testes para o código desenvolvido na primeira atividade ponderada.

<h3>Introdução à atividade</h3>

<p>Utilizando o simulador de dispositivos IoT desenvolvido na atividade passada e utilizando os conceitos de TDD vistos no decorrer da semana, implemente testes automatizados para validar o simulador.</p>

<p>Os testes devem abordar os seguintes requisitos:</p>

- Recebimento: garante que os dados enviados pelo simulador são recebidos pelo broker.
- Validação dos dados: garante que os dados enviados pelo simulador chegam sem alterações.
- Confirmação da taxa de disparo: garante que o simulador atende às especificações de taxa de disparo de mensagens dentro de uma margem de erro razoável.

<h3>Implementação</h3>

A fim de realizar a implementação da atividade, foram desenvolvidos um script que cumpre o papel de publisher ‒ <code>publisher.go</code> ‒ e um script que assume a função de subscriber ‒ <code>subscriber.go</code> ‒, cada um contido em seu respectivo diretório. Para garantir que os dados enviados pelo simulador fossem recebidos pelo broker, foi desenvolvido um script de teste <code>publisher_test.go</code>. O teste, ademais, realiza uma validação da taxa de disparo de mensagens, a fim de averiguar se a quantidade de mensagens recebidas em um determinado período de tempo está alinhada com o esperado.

A fim de executar o sistema, devem ser executados, em paralelo, os seguintes comandos:

*1.* Execução do broker local

```shell
mosquitto -c mosquitto.conf
```

*2.* Execução do publisher no diretório <code>publisher</code>

```shell
go run publisher.go
```

*3.* Execução do subscriber <code>subscriber</code>

```shell
go run subscriber.go
```

*4.* Execução do script de teste <code>publisher</code>

```shell
go test
```

<h3>Demonstração</h3>

Abaixo, encontra-se o vídeo demonstrativo gravado para a atividade. Nele, são seguidos os passos de execução descritos.