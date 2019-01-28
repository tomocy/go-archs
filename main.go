package main

import (
	"log"
	"net/http"

	"github.com/tomocy/archs/infra/web/server"
	"github.com/tomocy/archs/registry"
)

func main() {
	server := server.NewServer()
	registry := registry.NewRegistry()
	server.RegisterRoute(registry.NewHandler())
	log.Println("start listening and serving")
	if err := http.ListenAndServe(":5051", server); err != nil {
		panic(err)
	}
}
