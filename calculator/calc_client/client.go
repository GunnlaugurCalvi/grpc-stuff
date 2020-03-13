package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
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

	// n1, n2 := ValidateInput()
	c := calcpb.NewCalcServiceClient(cc)
	// doUnary(c, n1, n2)
	number := PrimeNumberInput()
	doServerStreaming(c, number)
}

func readInput(API_Structure string) string {
	reader := bufio.NewReader(os.Stdin)
	if API_Structure == "unary" {
		fmt.Println("Enter two numbers with seperated by whitespace: ")
	} else {
		fmt.Println("Enter a primenumber: ")
	}

	numbersString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("failed to read numbers from terminal input, %v", err)
	}

	return numbersString
}

func ValidateInput() (int64, int64) {
	numbersString := readInput("unary")
	nArray := strings.Split(numbersString, " ")
	n1, _ := strconv.ParseInt(nArray[0], 0, 64)
	tempn2 := strings.Replace(nArray[1], "\n", "", -1)
	n2, _ := strconv.ParseInt(tempn2, 0, 64)

	return n1, n2
}

func PrimeNumberInput() int64 {
	n := strings.Replace(readInput("stream"), "\n", "", -1)

	pnumber, err := strconv.ParseInt(n, 0, 64)
	if err != nil {
		log.Fatalf("failed to convert string to int, %v", err)
	}

	return pnumber
}

func doUnary(c calcpb.CalcServiceClient, n1, n2 int64) {
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

func doServerStreaming(c calcpb.CalcServiceClient, number int64) {
	fmt.Println("Performing prime decomposition")
	req := &calcpb.PrimeRequest{
		Num: number,
	}

	resp, err := c.PrimeDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to decomposit prime numbers, %v", err)
	}

	for {
		primeResponse, err := resp.Recv()
		if err == io.EOF {
			// has reached end of stream
			break
		}

		if err != nil {
			log.Fatalf("failed due to unexpected events %v", err)
		}

		fmt.Printf("number => %d\n", primeResponse.GetResult())
	}
}
