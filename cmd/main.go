package main

import (
	"log"

	gorestapi "github.com/KhanbalaRashidov/go-restapi"
)

func main() {
	srv := new(gorestapi.Server)

	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error ocurred  while  running  http  server: %s", err.Error())
	}
}
