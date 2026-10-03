package main

import (
	"log"

	"github.com/sboy99/exp.go/kitchen/services/kitchen/transports"
)

func main() {
	// Initialize and start the HTTP transport
	httpTransport := transports.NewHttpTransport(":8080")
	defer httpTransport.Stop()

	if err := httpTransport.Start(); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
