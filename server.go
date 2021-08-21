package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v \n", err)
	}

	grpcServer := grpc.NewServer()

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to serve gRPC on port 9000: %v \n", err)
	}
}

// Command used to create the chat.pd.go file
// protoc --go_out=plugins=grpc:./ chat.proto
