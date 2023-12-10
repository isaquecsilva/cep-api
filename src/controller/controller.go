package controller

import (
	"net/http"

	"github.com/isaquecsilva/cep-api/src/controller/utils"
	"github.com/isaquecsilva/cep-api/src/model"
	"github.com/isaquecsilva/cep-api/src/view"
	"github.com/labstack/echo"
)

type Controller struct {
	model.Model
}

func NewController(Model model.Model) *Controller {
	return &Controller{
		Model,
	}
}

func (c *Controller) GetCep(e echo.Context) error {
	cep := e.Param("cep")

	if err := utils.CepValidator(cep); err != nil {
		e.JSON(
			http.StatusBadRequest,
			map[string]any{
				"code":   http.StatusBadRequest,
				"errors": []string{
					err.Error(),
				},
			},
		)

		return nil
	}

	response := view.NewView(*c.Model.Execute(cep))
	e.JSON(response.Code, response)
	return nil
}
