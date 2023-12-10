package model

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedZipcodeQueryer struct {
	mock.Mock
}

func (mzq *MockedZipcodeQueryer) Execute(ctx context.Context, c chan ApiResponse, zipcode string) {
	select {
	case <-time.After(time.Second * time.Duration(rand.Intn(10))):
		c <- ApiResponse{
			StatusCode: 400,
			Errors: []error{
				fmt.Errorf("bad request").Error(),
			},
		}
	case <-ctx.Done():
	}
}

func Test_NewZipcodeQueryer(t *testing.T) {
	expected := &ZipcodeQueryer{
		gateways: []CepApi{
			&MockedZipcodeQueryer{},
		},
	}

	model := NewZipcodeQueryer(&MockedZipcodeQueryer{})
	assert.Equal(t, expected, model, "ZipcodeQueryer should be equal to specified struct")
}

func Test_Execute(t *testing.T) {
	model := NewZipcodeQueryer(&MockedZipcodeQueryer{}, &MockedZipcodeQueryer{})
	resp := model.Execute("fake_zipcode")

	expectation := &ApiResponse{
		StatusCode: http.StatusBadRequest,
		Errors: []error{
			fmt.Errorf("bad request").Error(),
		},
	}

	assert.Equal(t, expectation, resp, "ApiResponse")
}
