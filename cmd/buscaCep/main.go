package main

import (
	"fmt"
	"time"

	"github.com/biraneves/fc-go-multithreading/internal/model"
	"github.com/biraneves/fc-go-multithreading/internal/service"
)

func main() {
	ch1 := make(chan model.ResultDto)
	ch2 := make(chan model.ResultDto)
	cep := "89010025"

	go func() {
		dados, _, err := service.BuscaCep(model.ViaCep, cep)
		if err != nil {
			return
		}
		ch1 <- model.ResultDto{
			Endereco: model.EnderecoCompleto{
				Cep:        dados.Cep,
				Logradouro: dados.Logradouro,
				Bairro:     dados.Bairro,
				Cidade:     dados.Localidade,
				Estado:     dados.Uf,
			},
			Servico: "ViaCep",
		}
	}()

	go func() {
		_, dados, err := service.BuscaCep(model.CepBrasil, cep)
		if err != nil {
			return
		}
		ch2 <- model.ResultDto{
			Endereco: model.EnderecoCompleto{
				Cep:        dados.Cep,
				Logradouro: dados.Street,
				Bairro:     dados.Neighborhood,
				Cidade:     dados.City,
				Estado:     dados.State,
			},
			Servico: "CepBrasil",
		}
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg.Endereco, msg.Servico)

	case msg := <-ch2:
		fmt.Println(msg.Endereco, msg.Servico)

	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}
