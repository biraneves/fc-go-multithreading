# Projeto Multithreading em Go

Este repositório contém um projeto desenvolvido como parte de uma tarefa da pós-graduação Go Expert, da FC Tech, com foco em *multithreading* utilizando a linguagem Go.

## Descrição

O objetivo deste projeto é demonstrar a utilização de técnicas de multithreading em Go para realizar tarefas simultâneas de forma eficiente. O projeto inclui exemplos práticos e explicações detalhadas sobre como implementar e gerenciar threads em Go.

### Enunciado

> Neste desafio você terá que usar o que aprendemos com *multithreading* e *APIs* para buscar o resultado mais rápido entre duas *APIs* distintas.
> As duas requisições serão feitas simultaneamente para as seguintes *APIs*:
>
> `https://brasilapi.com.br/api/cep/v1/ + cep`
>
> `http://viacep.com.br/ws/" + cep + "/json/`
>
> Os requisitos para este desafio são:
> 
> - Acatar a *API* que entregar a resposta mais rápida e descartar a resposta mais lenta.
> - O resultado da *request* deverá ser exibido na *command line* com os dados do endereço, bem como qual *API* a enviou.
> - Limitar o tempo de resposta a 1 segundo. Caso contrário, o erro de *timeout* deve ser exibido.

## Tecnologias Utilizadas

- **Linguagem:** Go

## Como Executar

1. Clone o repositório para o seu ambiente local:
   ```bash
   git clone https://github.com/biraneves/fc-go-multithreading.git
   ```
2. Navegue até o diretório do projeto:
   ```bash
   cd fc-go-multithreading
   ```
3. Execute o projeto:
   ```bash
   go run cmd/buscaCep/main.go
   ```

## Contribuições

Contribuições são bem-vindas! Se você tiver sugestões ou melhorias, sinta-se à vontade para abrir
uma *issue* ou enviar um *pull request*.

