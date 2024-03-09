package main

import (
	"log"

	"github.com/shashankbhat10/Fetch-Backend-Assessment/api"
)

func main() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start()
	if err != nil {
		log.Fatal("error while starting server")
	}
}
