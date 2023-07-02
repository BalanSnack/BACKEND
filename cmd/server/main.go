package main

import (
	"BACKEND/config"
	"BACKEND/internals/app"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// configuration
	//config.Setup()
	config.Setup()
	// run server
	app.Run()
}
