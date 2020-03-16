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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("im calc client")
	cc, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial, %v", err)
	}
	reader := bufio.NewReader(os.Stdin)

	defer cc.Close()
	c := calcpb.NewCalcServiceClient(cc)
	fmt.Println(`choose API structure demo \n 
		Choose: \n 
		1) for Unary \n 
		2) for Server Streaming \n 
		3) for Client Streaming \n 
		4) for Bi-Direction Streaming \n
		5) for Error Unary Test`)

	option, err := reader.ReadString('\n')
	option = strings.Replace(option, "\n", "", -1)
	if err != nil {
		log.Fatalf("error getting input from user, %v", err)
	}
	switch option {
	case "1":
		n1, n2 := ValidateInput()
		doUnary(c, n1, n2)
	case "2":
		number := PrimeNumberInput()
		doServerStreaming(c, number)
	case "3":
		doClientStreaming(c)
	case "4":
		doBiDirectionStreaming(c)
	case "5":
		doErrorUnary(c)
	default:
		fmt.Println("Incorrect option")
	}
}

func readInput(API_Structure string) string {
	reader := bufio.NewReader(os.Stdin)

	switch API_Structure {
	case "unary":
		fmt.Println("Unary API RPC...")
		fmt.Println("Enter two numbers with seperated by whitespace: ")
	case "serverstream":
		fmt.Println("Server Streaming RPC..\nPerforming prime decomposition")
		fmt.Println("Enter a primenumber: ")
	case "clientstream":
		fmt.Println("Client Stream RPC...")
		fmt.Println("Enter as many numbers you want seperated by whitespace: ")
	case "bidirectional":
		fmt.Println("Bi Directional Streaming RPC...")
		fmt.Println("Will respond if number is larger than last maximum number(input numbers seperated by whitespace): ")
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
	fmt.Println("Multiplied together gives you", number)
}

func doClientStreaming(c calcpb.CalcServiceClient) {
	numberString := readInput("clientstream")

	numberString = strings.Replace(numberString, "\n", "", -1)
	numberList := strings.Split(numberString, " ")

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
		fmt.Printf("sending out %v\n", parseToInt)

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

func doBiDirectionStreaming(c calcpb.CalcServiceClient) {
	numberString := readInput("bidirectional")

	numberString = strings.Replace(numberString, "\n", "", -1)
	numberList := strings.Split(numberString, " ")

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Failed to create findmaximum, %v", err)
	}
	waitc := make(chan struct{})

	go func() {
		for _, n := range numberList {
			if n == " " || n == "" {
				continue
			}
			parseToInt, err := strconv.ParseInt(n, 0, 64)
			if err != nil {
				log.Fatalf("unable to parse string to int64, %v", err)
			}

			stream.Send(&calcpb.FindMaximumRequest{
				Number: parseToInt,
			})
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				fmt.Println(err)
				break
			}

			if err != nil {
				log.Fatalf("Failed to recieve from server stream, %v", err)
			}

			fmt.Printf("Current maximum: %d\n", res.GetMax())
		}

		close(waitc)
	}()

	<-waitc
}

func doErrorUnary(c calcpb.CalcServiceClient) {
	fmt.Println("starting to do sqrt unary RPC...")
	doErrorCall(c, 10)
	doErrorCall(c, -10)
}

func doErrorCall(c calcpb.CalcServiceClient, n int32) {
	resp, err := c.SquareRoot(context.Background(), &calcpb.SquareRootRequest{Number: n})

	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual error from gRPC (user error)
			fmt.Println("err message from server", respErr.Message())
			fmt.Println("err code from server", respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("we probably sent a negative number")
				return
			}
		} else {
			log.Fatalf("Big error calling square root: %v", err)
		}
	}

	fmt.Printf("Result of sqrt of %v: %v\n", n, resp.GetNumberRoot())
}
