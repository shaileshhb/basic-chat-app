package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

// SayHello is basic say hello function which takes message and return a new message.
func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {

	log.Printf("Received message body from client -> %s \n", message.Body)
	return &Message{Body: "Hello from server"}, nil
}
