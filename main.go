package main

import (
	"ecom/configuration"
	"ecom/server"
)

func main() {
	configurations := configuration.Configurations()
	port := configurations.AppPort

	server.StartServer(port)
}
