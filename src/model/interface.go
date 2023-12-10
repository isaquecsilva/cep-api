package model

import "context"

type Model interface {
	Execute(cep string) *ApiResponse
}

type CepApi interface {
	Execute(ctx context.Context, channel chan ApiResponse, zipCode string)
}