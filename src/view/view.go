package view

import "github.com/isaquecsilva/cep-api/src/model"

type Endereco = struct {
	Cep        string `json:"cep,omitempty"`
	Logradouro string `json:"logradouro,omitempty"`
	Bairro     string `json:"bairro,omitempty"`
	Localidade string `json:"localidade,omitempty"`
	UF         string `json:"uf,omitempty"`
}

type View struct {
	Code     int      `json:"code,omitempty"`
	Endereco          `json:"endereco,omitempty"`
	Errors   []string `json:"errors,omitempty"`
}

func NewView(ar model.ApiResponse) *View {
	if ar.Errors != nil {
		return &View{
			Code:   ar.StatusCode,
			Errors: ar.Errors,
		}
	}

	return &View{
		Code: ar.StatusCode,
		Endereco: Endereco{
			Cep:        ar.Body.Cep,
			Logradouro: ar.Body.Logradouro,
			Bairro:     ar.Body.Bairro,
			Localidade: ar.Body.Localidade,
			UF:         ar.Body.UF,
		},
	}
}
