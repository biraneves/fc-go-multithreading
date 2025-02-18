package model

type EnderecoCompleto struct {
	Cep        string
	Logradouro string
	Bairro     string
	Cidade     string
	Estado     string
}

type ResultDto struct {
	Endereco EnderecoCompleto
	Servico  string
}
