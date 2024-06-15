<h1 align="center"> 🌬️ Venture </h1>

<h1 align="center"> Somos segurança, velocidade e tecnologia. Somos Venture </h1>

<p align="center">
  <img src="https://i.imgur.com/yieDOSJ.png"/>
</p>

Este serviço é um consumer de um Broker Kafka (fila) que trabalha de maneira FIFO (first in, first out) e realiza o envio de email
junto com o AWS Simple Email Service, o AWS "SES". Caso exista alarmes / alertas de CPU, algo referente a isso. Alertas são recebidos através do SNS.

## ⚙️ Configurações

| Database                | Tabela        | Porta do Broker | Porta do Microserviço |
|-------------------------|---------------|-----------------|-----------------------|
| `venture_email_staging` | email_records | 9092            | 7788                  |
