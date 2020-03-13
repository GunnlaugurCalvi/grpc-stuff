package main

import (
	"context"
	"fmt"
	"io"
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
	// future implem let client choose API sturcture
	// switch case
	// doUnary(c)
	doServerStreaming(c)
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

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("starting to do a server streaming rpc...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Gulli",
			LastName:  "Calvii",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("err streaming many times %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}

		log.Printf("Response from GreetManyTimes: %s", msg.GetResult())
	}

}
