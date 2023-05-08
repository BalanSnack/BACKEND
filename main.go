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
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// run server
	app.Run(cfg)
}
