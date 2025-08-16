package main

import (
	"ecom/config"
	"ecom/server"
)

func main() {
	conf := config.Configurations()
	server.StartServer(conf)
}
