package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/Nextjingjing/go-god/12-grpc/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cred := insecure.NewCredentials()

	// create a gRPC client connection
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(cred))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// create a gRPC client
	client := pb.NewCalculatorServiceClient(conn)

	// call the Add method
	addResp, err := client.Add(ctx, &pb.CalculationRequest{Number1: 10, Number2: 5})
	if err != nil {
		panic(err)
	}
	fmt.Println(addResp.Result)

	// call the Subtract method
	subResp, err := client.Subtract(ctx, &pb.CalculationRequest{Number1: 10, Number2: 5})
	if err != nil {
		panic(err)
	}
	fmt.Println(subResp.Result)

	// call the Multiply method
	mulResp, err := client.Multiply(ctx, &pb.CalculationRequest{Number1: 10, Number2: 5})
	if err != nil {
		panic(err)
	}
	fmt.Println(mulResp.Result)

	// call the Divide method
	divResp, err := client.Divide(ctx, &pb.CalculationRequest{Number1: 10, Number2: 5})
	if err != nil {
		panic(err)
	}
	fmt.Println(divResp.Result)

	// call the Divide method with division by zero
	_, err = client.Divide(ctx, &pb.CalculationRequest{Number1: 10, Number2: 0})
	if err != nil {
		fmt.Println("Expected error:", err.Error())
	}

	// call the Average method
	stream, err := client.Average(ctx)
	if err != nil {
		panic(err)
	}
	numbers := []float64{10, 20, 30, 40, 50}
	for _, num := range numbers {
		if err := stream.Send(&pb.AverageRequest{Number: num}); err != nil {
			panic(err)
		}
	}
	// close the stream and receive the average response
	resp, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}
	fmt.Println("Average:", resp.Result)

	// call Multiplication table method
	mulTableStream, err := client.MultiplicationTable(ctx, &pb.MultiplicationTableRequest{Number: 5})
	if err != nil {
		panic(err)
	}
	for {
		resp, err := mulTableStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println("Multiplication Table Result:", resp.Result)
	}
}
