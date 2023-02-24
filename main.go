package main

import (
	"github.com/didnlie23/go-mvc/config"
	"github.com/didnlie23/go-mvc/internals/app"
)

func main() {
	// configuration
	config.Setup()
	// run server
	app.Run()
}
