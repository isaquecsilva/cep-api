package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/isaquecsilva/cep-api/src/controller"
)

func InitRouterAndServer(addr string, c *controller.Controller) (error) {
	e := echo.New()
	defer e.Close()
	e.Use(middleware.Logger())
	e.GET("/cep/:cep", c.GetCep)
	return e.Start(addr)
}