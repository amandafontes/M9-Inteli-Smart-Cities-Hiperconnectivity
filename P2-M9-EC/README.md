<h2>Atividade Prática | P2-M9-EC</h2>

<p>Diretório destinado à entrega da segunda prova do módulo 9 do curso de Engenharia da Computação.</p>

<h3>Introdução</h3>

O objetivo da atividade prática é a implementação de uma fila de eventos utilizando Apache Kafka, a fim de criar um amortecimento com persistência para que o sistema consiga absorver os momentos de pico sem perdas de dados. A ideia é permitir que o sistema de cidades inteligentes já existente aceite múltiplas requisições para armazenamento de dados de sensores ao mesmo tempo.

<h3>Implementação</h3>

Para a elaboração da atividade, foi, primeiramente, configurado o ambiente de desenvolvimento necessário para a execução de todas as tecnologias utilizadas, dentre as quais destaca-se a linguagem <code>.go</code> e o broker <code>Kafka</code>.

Posteriormente, foi necessário o desenvolvimento do código para a implementação da fila em Kafka, integrando-a ao producer e ao consumer da solução. O elemento em questão corresponde ao arquivo <code>kafka.go</code>.

No código, definimos a estrutura dos dados dos sensores responsáveis pela produção das informações coletadas, a função correspondente ao producer e a função correspondente ao consumer. Ambas são instanciadas na função principal, onde a fila em Kafka é devidamente implementada.