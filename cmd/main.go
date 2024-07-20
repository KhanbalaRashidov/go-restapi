package main

import (
	"github.com/KhanbalaRashidov/go-restapi/pkg/handler"
	"log"

	gorestapi "github.com/KhanbalaRashidov/go-restapi"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(gorestapi.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocurred  while  running  http  server: %s", err.Error())
	}
}
