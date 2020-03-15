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
	"time"

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

	// number := PrimeNumberInput()
	// doServerStreaming(c, number)

	doClientStreaming(c)
}

func readInput(API_Structure string) string {
	reader := bufio.NewReader(os.Stdin)

	switch API_Structure {
	case "unary":
		fmt.Println("Enter two numbers with seperated by whitespace: ")
	case "serverstream":
		fmt.Println("Enter a primenumber: ")
	case "clientstream":
		fmt.Println("Enter as many numbers you want seperated by comma: ")
	default:
		fmt.Println("weird error")
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
	n := strings.Replace(readInput("serverstream"), "\n", "", -1)

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

func doClientStreaming(c calcpb.CalcServiceClient) {
	fmt.Println("Client Stream RPC...")

	numberString := readInput("clientstream")

	numberString = strings.Replace(numberString, "\n", "", -1)
	numberList := strings.Split(numberString, " ")

	fmt.Println("list of numbers", numberList, len(numberList))

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Failed to compute average: %v", err)
	}

	for _, n := range numberList {
		if n == " " || n == "" {
			continue
		}
		parseToInt, err := strconv.ParseInt(n, 0, 64)
		if err != nil {
			log.Fatalf("unable to parse string to int64, %v", err)
		}
		fmt.Printf("sending out %v, %T\n", parseToInt, parseToInt)

		stream.Send(&calcpb.ComputeAvgRequest{
			Number: parseToInt,
		})

		time.Sleep(100 * time.Millisecond)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error closing stream, %v", err)
	}

	fmt.Printf("Server response %v", resp.GetAvgResult())
}
