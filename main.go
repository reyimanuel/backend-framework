package main

import (
	"backend/config"
	"backend/internal/server"
	"backend/pkg/token"
)

func main() {
	config.Load()
	token.Load()
	server.Run()
}

//  Package main is the entry point for the application. It loads the configuration and starts the server.
