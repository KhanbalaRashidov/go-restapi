package main

import (
	"github.com/KhanbalaRashidov/go-restapi/pkg/handler"
	"github.com/KhanbalaRashidov/go-restapi/pkg/repository"
	"github.com/KhanbalaRashidov/go-restapi/pkg/service"
	"log"

	gorestapi "github.com/KhanbalaRashidov/go-restapi"
)

func main() {
	repos := repository.NeRepository()
	services := service.NeService(repos)
	handlers := handler.NewHandler(services)

	srv := new(gorestapi.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocurred  while  running  http  server: %s", err.Error())
	}
}
