package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/biraneves/fc-go-multithreading/internal/model"
)

func BuscaCep(servico int, cep string) (*model.ViaCepApi, *model.CepBrasilApi, error) {
	url, err := urlServico(servico, cep)
	if err != nil {
		return nil, nil, err
	}

	data, err := buscaDados(url)
	if err != nil {
		return nil, nil, err
	}

	switch servico {
	case model.ViaCep:
		var viaCepData model.ViaCepApi
		if err := json.Unmarshal(data, &viaCepData); err != nil {
			return nil, nil, err
		}
		return &viaCepData, nil, nil

	case model.CepBrasil:
		var brasilApiData model.CepBrasilApi
		if err := json.Unmarshal(data, &brasilApiData); err != nil {
			return nil, nil, err
		}
		return nil, &brasilApiData, nil
	}

	return nil, nil, nil
}

func urlServico(servico int, cep string) (string, error) {
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

func buscaDados(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
