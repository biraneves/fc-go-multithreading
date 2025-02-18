package service

import (
	"fmt"
	"strings"

	"github.com/biraneves/fc-go-multithreading/internal/model"
)

func BuscaCep(servico int, cep string) (*model.ViaCepApi, *model.CepBrasilApi, error) {
	return nil, nil, nil
}

func UrlServico(servico int, cep string) (string, error) {
	var url string

	switch servico {
	case model.CepBrasil:
		url = "https://brasilapi.com.br/api/cep/v1/{cep}"
	case model.ViaCep:
		url = "http://viacep.com.br/ws/{cep}/json/"
	default:
		url = ""
	}

	if url == "" {
		return "", fmt.Errorf("serviço desconhecido")
	}

	finalUrl := strings.ReplaceAll(url, "{cep}", cep)

	return finalUrl, nil
}
