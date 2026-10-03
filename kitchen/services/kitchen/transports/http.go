package transports

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sboy99/exp.go/kitchen/services/kitchen/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HttpTransport struct {
	router     *mux.Router
	grpcClient *grpc.ClientConn
	address    string
}

func NewGrpcClient(address string) (*grpc.ClientConn, error) {
	return grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func NewHttpTransport(address string) *HttpTransport {
	grpcClient, err := NewGrpcClient("localhost:9090") // Assuming the gRPC server is running on localhost:50051
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	return &HttpTransport{
		router:     mux.NewRouter(),
		grpcClient: grpcClient,
		address:    address,
	}
}

func (h *HttpTransport) Start() error {
	// Start the HTTP server here
	server := &http.Server{
		Addr:         h.address,
		Handler:      h.router,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
	}
	h.RegisterRoutes()
	log.Printf("Starting HTTP server on %s", h.address)
	return server.ListenAndServe()
}

func (h *HttpTransport) Stop() {
	// Implement graceful shutdown logic here if needed
	h.grpcClient.Close()
	log.Printf("Stopping HTTP server on %s", h.address)
}

func (h *HttpTransport) RegisterRoutes() {
	orderHandler := handler.NewOrderHttpHandler(h.grpcClient)
	h.router.HandleFunc("/orders", orderHandler.CreateOrderHandler).Methods("POST")
	h.router.HandleFunc("/orders/{id}", orderHandler.GetOrderHandler).Methods("GET")
	h.router.HandleFunc("/orders", orderHandler.ListOrdersHandler).Methods("GET")
}
