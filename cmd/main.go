package main

import (
	"MoneyService"
	"MoneyService/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(MoneyService.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured white running http server: %s", err.Error())
	}
}
