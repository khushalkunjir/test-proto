package chat

import (
	fmt "fmt"
	"io"
	"log"

	"github.com/khushalkunjir/test-proto/chat"
	"golang.org/x/net/context"
)

type Server struct {
}
type Server1 struct {
}

type StramService struct {
	chat.UnimplementedStramServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}

func (s *Server1) SendParameters(ctx context.Context, instruction *Instruction) (*Instruction, error) {
	log.Printf("Received Instruction Ip address from client: %s", instruction.IpAddress)
	log.Printf("Received Instruction port from client: %s", instruction.Port)
	log.Printf("Received Instruction name from client: %s", instruction.Name)
	return &Instruction{IpAddress: "10.10.10.1", Port: "9000", Name: "Node1"}, nil
}

func (s *StramService) CalculateBeatsPerMinute(stream chat.StramService_CalculateBeatsPerMinuteServer) error {
	var count, total uint32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil
		}
		total += req.GetValue()

		fmt.Println("Received", req.GetValue())

		count++
		if count%5 == 0 {
			fmt.Println("Total:", total, "Sending:", float32(total)/5)
			if err := stream.Send(&chat.CalculateBeatsPerMinuteResponse{
				Average: float32(total) / 5,
			}); err != nil {
				return nil
			}
			total = 0

		}

	}
}

