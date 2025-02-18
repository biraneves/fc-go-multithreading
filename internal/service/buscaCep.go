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
	url, err := UrlServico(servico, cep)
	if err != nil {
		return nil, nil, err
	}

	if servico == model.ViaCep {
		res, err := http.Get(url)
		if err != nil {
			return nil, nil, err
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, nil, err
		}

		var viaCepData model.ViaCepApi
		err = json.Unmarshal(data, &viaCepData)
		if err != nil {
			return nil, nil, err
		}

		return &viaCepData, nil, nil
	}

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
