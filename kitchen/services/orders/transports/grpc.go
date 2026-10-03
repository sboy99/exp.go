package transports

import (
	"net"

	"log"

	"github.com/sboy99/exp.go/kitchen/services/orders/handler"
	"github.com/sboy99/exp.go/kitchen/services/orders/service"
	"google.golang.org/grpc"
)

type GrpcTransport struct {
	address string
	server  *grpc.Server
}

func NewGrpcTransport(address string) *GrpcTransport {
	return &GrpcTransport{
		server:  grpc.NewServer(),
		address: address,
	}
}

func (g *GrpcTransport) Start() error {
	listener, err := g.createListener()
	if err != nil {
		return err
	}

	// Register handlers here
	orderSvc := service.NewOrderService()
	handler.NewOrderGRPCHandler(g.server, orderSvc)

	log.Printf("Starting gRPC server on %s", g.address)
	return g.server.Serve(listener)
}

func (g *GrpcTransport) Stop() error {
	g.server.GracefulStop()
	return nil
}

func (g *GrpcTransport) createListener() (net.Listener, error) {
	listener, err := net.Listen("tcp", g.address)
	return listener, err
}
