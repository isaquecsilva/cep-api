package opencep

import (
	"net/http"
	"time"
	"context"
	"fmt"
	"io"
	"encoding/json"
	"strconv"

	"github.com/isaquecsilva/cep-api/src/model"
)

type OpenCep struct {
	timeout  time.Duration
	endpoint string
}

type OpenCepResponse struct {
	Cep         string `json:"cep,omitempty"`
	Logradouro  string `json:"logradouro,omitempty"`
	Complemento string `json:"complemento,omitempty"`
	Bairro      string `json:"bairro,omitempty"`
	Localidade  string `json:"localidade,omitempty"`
	UF          string `json:"uf,omitempty"`
	IBGE        string `json:"ibge,omitempty"`
}

func NewOpenCep() *OpenCep {
	return &OpenCep{
		endpoint: "https://opencep.com/v1/%s",
	}
}

func (oc *OpenCep) Execute(ctx context.Context, channel chan model.ApiResponse, zipCode string) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(oc.endpoint, zipCode), nil)

	if err != nil {
		channel <- model.ApiResponse{
			StatusCode: http.StatusInternalServerError,
			Errors: []string{
				err.Error(),
			},
		}

		return
	} else {


		// Getting request from http package
		client := http.DefaultClient
		
		// Doing the request
		res, err := client.Do(req)

		// Checking for errors.
		// Whether there is someone, we return a ApiResponse
		// filled with the returned error
		if err != nil {
			channel <- model.ApiResponse{
				StatusCode: http.StatusInternalServerError,
				Errors: []string{err.Error()},
			}

			return
		}
		// Defering body close
		defer res.Body.Close()

		// Reading response of our done
		// request
		buf, err := io.ReadAll(res.Body)
		if err != nil {
			channel <- model.ApiResponse{
				StatusCode: http.StatusInternalServerError,
				Errors: []string{
					err.Error(),
				},
			}
			return
		}

		// Unmarshaling the api response
		var ocr OpenCepResponse
		err = json.Unmarshal(buf, &ocr)

		// Checking for unmarshaling errors
		if err != nil {
			channel <- model.ApiResponse{
				StatusCode: http.StatusInternalServerError,
				Errors: []string{
					err.Error(),
				},
			}

		} else {
			ibgeInt64, _ := strconv.ParseInt(ocr.IBGE, 10, 64)

			channel <- model.ApiResponse{
				StatusCode: res.StatusCode,
				Body: model.Body{
					Cep:        ocr.Cep,
					Logradouro: ocr.Logradouro,
					Bairro:     ocr.Bairro,
					Localidade: ocr.Localidade,
					UF:         ocr.UF,
					Ibge:       ibgeInt64,
				},
			}
		}

	}
}