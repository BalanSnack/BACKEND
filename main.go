package main

import (
	"github.com/BalanSnack/BACKEND/config"
	"github.com/BalanSnack/BACKEND/internals/app"
)

func main() {
	// configuration
	config.Setup()
	// run server
	app.Run()
}
