package main

import (
	"github.com/YghoRibas/Dito-hexAPI/src/port"
)

func main() {
	server := port.Server{}

	server.Run()
}