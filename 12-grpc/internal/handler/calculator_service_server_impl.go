package handler

import (
	"context"
	"errors"
	"io"

	"github.com/Nextjingjing/go-god/12-grpc/internal/pb"
	"google.golang.org/grpc"
)

type calculatorServiceServerImpl struct {
	pb.UnimplementedCalculatorServiceServer
}

func NewCalculatorServiceServer() pb.CalculatorServiceServer {
	return &calculatorServiceServerImpl{}
}

func (s *calculatorServiceServerImpl) Add(ctx context.Context, req *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	result := req.Number1 + req.Number2
	return &pb.CalculationResponse{
		Result: result,
	}, nil
}

func (s *calculatorServiceServerImpl) Subtract(ctx context.Context, req *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	result := req.Number1 - req.Number2
	return &pb.CalculationResponse{
		Result: result,
	}, nil
}

func (s *calculatorServiceServerImpl) Multiply(ctx context.Context, req *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	result := req.Number1 * req.Number2
	return &pb.CalculationResponse{
		Result: result,
	}, nil
}

func (s *calculatorServiceServerImpl) Divide(ctx context.Context, req *pb.CalculationRequest) (*pb.CalculationResponse, error) {
	if req.Number2 == 0 {
		return nil, errors.New("Error: Division by zero")
	}
	result := req.Number1 / req.Number2
	return &pb.CalculationResponse{
		Result: result,
	}, nil
}

func (s *calculatorServiceServerImpl) Average(stream grpc.ClientStreamingServer[pb.AverageRequest, pb.CalculationResponse]) error {
	sum := 0.0
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		sum += req.Number
		count++
	}
	if count == 0 {
		return errors.New("Error: No numbers provided")
	}
	average := sum / float64(count)
	return stream.SendAndClose(&pb.CalculationResponse{
		Result: average,
	})
}

func (s *calculatorServiceServerImpl) MultiplicationTable(req *pb.MultiplicationTableRequest, stream grpc.ServerStreamingServer[pb.CalculationResponse]) error {
	for i := 1; i <= 12; i++ {
		result := req.Number * float64(i)
		err := stream.Send(&pb.CalculationResponse{
			Result: result,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
