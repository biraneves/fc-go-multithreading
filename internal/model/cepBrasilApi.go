package model

type CepBrasilApi struct {
	Cep          string `json:"cep,omitempty"`
	State        string `json:"state,omitempty"`
	City         string `json:"city,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
	Street       string `json:"street,omitempty"`
	// Service      string `json:"service,omitempty"`
}
