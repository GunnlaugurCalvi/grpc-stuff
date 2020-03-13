package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gunnlaugurcalvi/grpc-stuff/calculator/calcpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("im calc client")
	cc, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial, %v", err)
	}

	defer cc.Close()

	n1, n2 := readAndValidateInput()
	c := calcpb.NewSumServiceClient(cc)
	doUnary(c, n1, n2)
}

func readAndValidateInput() (int64, int64) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter two numbers with seperated by whitespace: ")
	numbersString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("failed to read numbers from terminal input, %v", err)
	}

	nArray := strings.Split(numbersString, " ")
	n1, _ := strconv.ParseInt(nArray[0], 0, 64)
	tempn2 := strings.Replace(nArray[1], "\n", "", -1)
	n2, _ := strconv.ParseInt(tempn2, 0, 64)

	return n1, n2
}

func doUnary(c calcpb.SumServiceClient, n1, n2 int64) {
	fmt.Println("client said hi")
	req := &calcpb.SumRequest{
		SumResult: &calcpb.SumNumbers{
			Num_1: n1,
			Num_2: n2,
		},
	}

	resp, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to perform sum %v", err)
	}

	fmt.Println("the sum of these two numbers is: ", resp.Result)
}
