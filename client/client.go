package main

import (
	"log"

	"github.com/shaileshhb/basic-chat-app/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection failed: %v \n", err)
	}
	defer conn.Close()

	chatClient := chat.NewBasicChatServiceClient(conn)

	response, err := chatClient.SayHello(context.Background(), &chat.Message{
		Body: "Hello From client",
	})
	if err != nil {
		log.Fatalf("Error while calling SayHello: %v \n", err)
	}

	log.Println("Response from server ->", response.Body)
}
