package main

import (
	"log"
	"net"

	"github.com/shaileshhb/basic-chat-app/chat"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v \n", err)
	}

	log.Println("Server started at port :9000")

	chatServer := chat.Server{}
	grpcServer := grpc.NewServer()

	chat.RegisterBasicChatServiceServer(grpcServer, &chatServer)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to serve gRPC on port 9000: %v \n", err)
	}
}

// Command used to create the chat.pd.go file
// protoc --go_out=plugins=grpc:./ chat.proto
