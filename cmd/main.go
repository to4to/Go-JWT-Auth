package main

import (
	"log"

	"github.com/to4to/Go-JWT-Auth/cmd/api"
)

func main() {

	server := api.NewApiServer(":8080", nil)

	if err := server.Run(); err != nil {

		log.Fatal(err)

	}
}
