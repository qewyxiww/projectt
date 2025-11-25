package main

import (
	"log"
	"os"

	"github.com/qewyxiww/projectt/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "MORSE-CONVERTER: ", log.LstdFlags|log.Lshortfile)

	srv := server.New(logger)

	if err := srv.Start(); err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}
