package main

import (
	"github.com/BalanSnack/BACKEND/config"
	"github.com/BalanSnack/BACKEND/internals/app"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// configuration
	config.Setup()
	// run server
	app.Run()
}
