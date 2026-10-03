package main

import (
	"log"

	"github.com/sboy99/exp.go/kitchen/services/orders/transports"
)

func main() {
	grpcTransport := transports.NewGrpcTransport(":9090")
	defer grpcTransport.Stop()

	if err := grpcTransport.Start(); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}

}
