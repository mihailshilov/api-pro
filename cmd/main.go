package main

import (
	"log"

	apipro "github.com/mihailshilov/api-pro"
	"github.com/mihailshilov/api-pro/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(apipro.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
