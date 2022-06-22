package main

import (
	"github.com/kaliayev-proton/banking-go-hex/app"
	"github.com/kaliayev-proton/banking-go-hex/logger"
)

func main() {
	// log.Println("Starting our application...")
	logger.Info("Starting our application...")
	app.Start()
}
