package main

import (
	"backend/config"
	"backend/internal/server"
)

func main() {
	config.Load()
	server.Run()
}

//  Package main is the entry point for the application. It loads the configuration and starts the server.
