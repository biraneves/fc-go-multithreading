package service

import (
	"encoding/json"
	"testing"

	"github.com/biraneves/fc-go-multithreading/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestUrlServico(t *testing.T) {
	cep := "01001000"

	servico := model.CepBrasil
	expectedUrlCepBrasil := "https://brasilapi.com.br/api/cep/v1/01001000"
	actualUrlCepBrasil, err := urlServico(servico, cep)
	assert.Nil(t, err)
	assert.NotEmpty(t, actualUrlCepBrasil)
	assert.Equal(t, expectedUrlCepBrasil, actualUrlCepBrasil)

	servico = model.ViaCep
	expectedUrlViaCep := "http://viacep.com.br/ws/01001000/json/"
	actualUrlViaCep, err := urlServico(servico, cep)
	assert.Nil(t, err)
	assert.NotEmpty(t, actualUrlViaCep)
	assert.Equal(t, expectedUrlViaCep, actualUrlViaCep)
}

func TestBuscaCep(t *testing.T) {
	cep := "01001000"

	servico := model.ViaCep
	expectedResViaCepApi := &model.ViaCepApi{
		Cep:         "01001-000",
		Logradouro:  "Praça da Sé",
		Complemento: "lado ímpar",
		Unidade:     "",
		Bairro:      "Sé",
		Localidade:  "São Paulo",
		Uf:          "SP",
		Estado:      "São Paulo",
		Regiao:      "Sudeste",
		Ibge:        "3550308",
		Gia:         "1004",
		Ddd:         "11",
		Siafi:       "7107",
	}

	actualResViaCepApi, actualResBrasilApi, actualError := BuscaCep(servico, cep)
	assert.Nil(t, actualError)
	assert.Nil(t, actualResBrasilApi)
	assert.NotNil(t, actualResViaCepApi)
	assert.Equal(t, expectedResViaCepApi, actualResViaCepApi)

	servico = model.CepBrasil
	expectedResCepBrasilApi := &model.CepBrasilApi{
		Cep:          "01001000",
		State:        "SP",
		City:         "São Paulo",
		Neighborhood: "Sé",
		Street:       "Praça da Sé",
	}

	actualResViaCepApi, actualResBrasilApi, actualError = BuscaCep(servico, cep)
	assert.Nil(t, actualError)
	assert.Nil(t, actualResViaCepApi)
	assert.NotNil(t, actualResBrasilApi)
	assert.Equal(t, expectedResCepBrasilApi, actualResBrasilApi)
}

func TestBuscaDados(t *testing.T) {
	urlViaCep := "http://viacep.com.br/ws/01001000/json/"
	var resViaCepApi model.ViaCepApi
	expectedResViaCepApi := &model.ViaCepApi{
		Cep:         "01001-000",
		Logradouro:  "Praça da Sé",
		Complemento: "lado ímpar",
		Unidade:     "",
		Bairro:      "Sé",
		Localidade:  "São Paulo",
		Uf:          "SP",
		Estado:      "São Paulo",
		Regiao:      "Sudeste",
		Ibge:        "3550308",
		Gia:         "1004",
		Ddd:         "11",
		Siafi:       "7107",
	}
	actualResViaCepApi, errResViaCepApi := buscaDados(urlViaCep)
	errJson := json.Unmarshal(actualResViaCepApi, &resViaCepApi)
	assert.Nil(t, errResViaCepApi)
	assert.NotNil(t, actualResViaCepApi)
	assert.Nil(t, errJson)
	assert.NotNil(t, resViaCepApi)
	assert.Equal(t, expectedResViaCepApi, &resViaCepApi)

	urlCepBrasilApi := "https://brasilapi.com.br/api/cep/v1/01001000"
	var resCepBrasilApi model.CepBrasilApi
	expectedResCepBrasilApi := &model.CepBrasilApi{
		Cep:          "01001000",
		State:        "SP",
		City:         "São Paulo",
		Neighborhood: "Sé",
		Street:       "Praça da Sé",
	}
	actualResCepBrasilApi, errResCepBrasilApi := buscaDados(urlCepBrasilApi)
	errJson = json.Unmarshal(actualResCepBrasilApi, &resCepBrasilApi)
	assert.Nil(t, errResCepBrasilApi)
	assert.NotNil(t, actualResCepBrasilApi)
	assert.Nil(t, errJson)
	assert.NotNil(t, resCepBrasilApi)
	assert.Equal(t, expectedResCepBrasilApi, &resCepBrasilApi)

}
