package view

import (
	"testing"

	"github.com/isaquecsilva/cep-api/src/model"
	"github.com/stretchr/testify/assert"
)

func Test_NewView(t *testing.T) {
	resp := model.ApiResponse{
		StatusCode: 200,
		Body: struct {
			Cep         string
			Logradouro  string
			Complemento string
			Bairro      string
			Localidade  string
			UF          string
			Ibge        int64
		}{"00000000","fake_logradouro","","bairro","cidade","MG",12345},
		Errors: nil,
	}

	v := NewView(resp)

	var expected *View = &View{
		Code: 200,
		Endereco: Endereco{
			Cep:        "00000000",
			Logradouro: "fake_logradouro",
			Bairro:     "bairro",
			Localidade: "cidade",
			UF:         "MG",
		},
		Errors: nil,
	}

	assert.Equal(t, expected, v, "View generated struct is not equal to expected")
}
