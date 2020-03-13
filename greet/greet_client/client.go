package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gunnlaugurcalvi/grpc-stuff/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("im client")

	cc, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial, %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("starting unary rpc")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Gulli",
			LastName:  "Calvi",
		},
	}

	resp, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to greet %v", err)
	}

	fmt.Printf("Response from greet: %v", resp.Result)
}
