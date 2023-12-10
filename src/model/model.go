package model

import (
	"context"
)

type Body struct {
	Cep         string
	Logradouro  string
	Complemento string
	Bairro      string
	Localidade  string
	UF          string
	Ibge        int64
}

type ApiResponse struct {
	StatusCode int
	Body
	Errors []string
}

type ZipcodeQueryer struct {
	gateways []CepApi
}

func NewZipcodeQueryer(gateways ...CepApi) *ZipcodeQueryer {
	return &ZipcodeQueryer{
		gateways,
	}
}

func (zq *ZipcodeQueryer) Execute(zipcode string) *ApiResponse {
	channel := make(chan ApiResponse)
	ctx, cancel := context.WithCancel(context.Background())

	// Making simultaneous requests to our cep
	// services.
	for i := 0; i < len(zq.gateways); i++ {
		go zq.gateways[i].Execute(ctx, channel, zipcode)
	}

	// The first service that respond us, we cancel
	// the others.
	select {
	case response := <-channel:
		cancel()
		return &response
	}
}
