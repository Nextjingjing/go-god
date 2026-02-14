package handler

import (
	"context"
	"errors"

	"github.com/Nextjingjing/go-god/12-grpc/internal/pb"
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
