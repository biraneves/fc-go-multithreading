package service

import (
	"testing"

	"github.com/biraneves/fc-go-multithreading/internal/model"
	"github.com/biraneves/fc-go-multithreading/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestUrlServico(t *testing.T) {
	cep := "01001000"

	servico := model.CepBrasil
	expectedUrlCepBrasil := "https://brasilapi.com.br/api/cep/v1/01001000"
	actualUrlCepBrasil, err := service.UrlServico(servico, cep)
	assert.Nil(t, err)
	assert.NotEmpty(t, actualUrlCepBrasil)
	assert.Equal(t, expectedUrlCepBrasil, actualUrlCepBrasil)

	servico = model.ViaCep
	expectedUrlViaCep := "http://viacep.com.br/ws/01001000/json/"
	actualUrlViaCep, err := service.UrlServico(servico, cep)
	assert.Nil(t, err)
	assert.NotEmpty(t, actualUrlViaCep)
	assert.Equal(t, expectedUrlViaCep, actualUrlViaCep)
}
