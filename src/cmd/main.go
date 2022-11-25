package main

import (
	"github.com/YghoRibas/Dito-hexAPI/src/app"
	"github.com/YghoRibas/Dito-hexAPI/src/port"
)

func main() {

	app := app.Application{}
	server := port.Server{App: app}

	server.Run()
}
