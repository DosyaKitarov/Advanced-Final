package main

import (
	"log"
	"net"

	"library/pkg/grpc"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer(":50051", ":50052")
	grpc.RegisterLibraryServiceServer(server, server)

	log.Println("Starting Library Service on port 50053...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
