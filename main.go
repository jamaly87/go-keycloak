package main

import (
	"keycloak/controller"
	"keycloak/keycloak"
	"log"
)

const ServerAddress string = "0.0.0.0:1010"

func main() {
	server := controller.NewServer(keycloak.NewKeycloak())
	err := server.Start(ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
