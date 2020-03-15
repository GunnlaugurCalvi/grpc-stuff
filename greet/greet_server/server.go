package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/gunnlaugurcalvi/grpc-stuff/greet/greetpb"

	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	result := fmt.Sprintf("Hello %s %s", firstName, lastName)

	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func (s *Server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Println("Greet stream API invoked")
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := fmt.Sprintf("heyyooo %s req nr %d", firstName, i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}

		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func (s *Server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Println("LongGreet func was invoked with a a streaming request")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("failed to recv %v", err)
		}

		fName := req.GetGreeting().GetFirstName()

		result += fmt.Sprintf(" Hello %s!", fName)

	}
}

func main() {
	fmt.Println("Greet Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("failed to listen, %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
