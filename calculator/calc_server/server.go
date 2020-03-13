package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/gunnlaugurcalvi/grpc-stuff/calculator/calcpb"
	"google.golang.org/grpc"
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
			time.Sleep(1000 * time.Millisecond)
			num /= k
		} else {
			k++
		}
	}

	return nil
}

func main() {
	fmt.Println("init server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen, %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalcServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
