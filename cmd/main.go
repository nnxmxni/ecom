package main

import (
	"github.com/nnxmxni/ecom/cmd/api"
	"log"
)

func main() {

	server := api.NewAPIServer(":8000", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
