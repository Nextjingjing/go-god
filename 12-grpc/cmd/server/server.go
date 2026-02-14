package main

import (
	"log"
	"net"

	"github.com/Nextjingjing/go-god/12-grpc/internal/handler"
	"github.com/Nextjingjing/go-god/12-grpc/internal/pb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(grpcServer, handler.NewCalculatorServiceServer())
	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
