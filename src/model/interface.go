package model

import "context"

type CepApi interface {
	Execute(ctx context.Context, channel chan ApiResponse, zipCode string)
}
