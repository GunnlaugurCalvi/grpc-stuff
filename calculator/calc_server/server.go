package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"

	"github.com/gunnlaugurcalvi/grpc-stuff/calculator/calcpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// Server is server
type Server struct{}

// Sum takes in two numbers and returns the sum of them
func (s *Server) Sum(ctx context.Context, req *calcpb.SumRequest) (*calcpb.SumResponse, error) {
	fmt.Println("Executing sum of the numbers")

	num1 := req.GetSumResult().GetNum_1()
	num2 := req.GetSumResult().GetNum_2()
	sum := num1 + num2

	result := &calcpb.SumResponse{
		Result: sum,
	}

	return result, nil
}

// PrimeDecomposition decomposits the number that is given
func (s *Server) PrimeDecomposition(req *calcpb.PrimeRequest, stream calcpb.CalcService_PrimeDecompositionServer) error {
	fmt.Println("Prime Decomposition")
	num := req.GetNum()
	k := int64(2)

	for num > 1 {
		if num%k == 0 {
			res := &calcpb.PrimeResponse{
				Result: k,
			}
			stream.Send(res)
			num /= k
		} else {
			k++
		}
	}

	return nil
}

// ComputeAverage computes average of numbers that client streams
func (s *Server) ComputeAverage(stream calcpb.CalcService_ComputeAverageServer) error {
	fmt.Println("Computing average...")
	var number int64
	counter := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have reached end of stream
			return stream.SendAndClose(&calcpb.ComputeAvgResponse{
				AvgResult: float64(number) / float64(counter),
			})
		}

		if err != nil {
			log.Fatalf("Stream failure %v", err)
		}

		counter++
		number += req.GetNumber()
	}
}

// FindMaximum sends back to client if current number is bigger than last one
func (s *Server) FindMaximum(stream calcpb.CalcService_FindMaximumServer) error {
	fmt.Println("Finding maximum from client stream...")
	lastNumber := int64(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("failed to recieve from client stream, %v", err)
		}

		currentNumber := req.GetNumber()
		if currentNumber > lastNumber {
			err = stream.Send(&calcpb.FindMaximumResponse{
				Max: currentNumber,
			})

			if err != nil {
				return err
			}

			lastNumber = currentNumber
		}
	}
}

func (s *Server) SquareRoot(ctx context.Context, req *calcpb.SquareRootRequest) (*calcpb.SquareRootResponse, error) {
	fmt.Println("Recieved SquareRoot RPC")

	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Recieved negative number: %v", number))
	}

	return &calcpb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
}

func main() {
	fmt.Println("init server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen, %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalcServiceServer(s, &Server{})

	// Register reflection service on gRPC server
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
