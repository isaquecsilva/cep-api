package opencep

import (
	"net/http"

	"github.com/isaquecsilva/cep-api/src/model"
)

type OpenCep struct {
	timeout  time.Duration
	endpoint string
}

func NewOpenCep() *OpenCep {
	return &OpenCep{
		endpoint: "opencep.com/v1/%s",
	}
}

func (oc *OpenCep) Execute(ctx context.Context, channel chan model.ApiResponse, zipCode string) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(oc.endpoint, zipCode), nil)

	if err != nil {
		channel <- model.ApiResponse{
			StatusCode: http.StatusInternalServerError,
			Errors: []error{
				err,
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
				Errors: []error{err},
			}

			return
		}
		// Defering body close
		defer res.Body.Close()

		// Reading response of our done
		// request
		buf, err := io.ReadAll(res)
		if err != nil {
			channel <- model.ApiResponse{
				StatusCode: http.StatusInternalServerError,
				Errors: []error{
					err,
				},
			}
			return
		}

		// Unmarshaling the api response
		var ar *model.ApiResponse
		err := json.Unmarshal(buf, ar)

		// Checking for unmarshaling errors
		if err != nil {
			channel <- model.ApiResponse{
				StatusCode: http.StatusInternalServerError,
				Errors: []error{
					err,
				},
			}

		} else {

			channel <- *ar
		}

	}
}