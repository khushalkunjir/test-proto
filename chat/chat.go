package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf(ctx)
	log.Printf("Received message body from client: %s", message.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}

func (s *Server) SendParameters(ctx context.Context, instruction Instruction) (*Instruction, error) {
	log.Printf(ctx)
	log.Printf("Received Instruction Ip address from client: %s", instruction.IpAddress)
	log.Printf("Received Instruction port from client: %s", instruction.Port)
	log.Printf("Received Instruction name from client: %s", instruction.Name)
	return &Instruction{IpAddress: "10.10.10.1", Port: "9000", Name: "Node1"}, nil
}
