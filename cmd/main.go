package main

import (
	"log"

	"github.com/isaquecsilva/cep-api/src/model"
	"github.com/isaquecsilva/cep-api/src/controller"
	"github.com/isaquecsilva/cep-api/src/gateway/opencep"
	"github.com/isaquecsilva/cep-api/src/routes"
)

func main() {
	opencep := opencep.NewOpenCep()
	service := model.NewZipcodeQueryer(opencep)
	control := controller.NewController(service)
	log.Fatal(routes.InitRouterAndServer("localhost:8080", control))
}