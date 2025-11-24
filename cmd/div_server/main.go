package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	divpb "notask/op-coordinator/api/divpb"

	"google.golang.org/grpc"
)

type DivServer struct {
	divpb.UnimplementedDivServiceServer
}

func (s *DivServer) Compute(ctx context.Context, req *divpb.OperationRequest) (*divpb.OperationResponse, error) {
	if req.B == 0 {
		return nil, fmt.Errorf("division by zero")
	}

	return &divpb.OperationResponse{
		Result: req.A / req.B,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5004")
	if err != nil {
		log.Fatalf("Error opening port 5004: %v", err)
	}

	grpcServer := grpc.NewServer()
	divpb.RegisterDivServiceServer(grpcServer, &DivServer{})

	go func() {
		log.Println("DivService running on :5004")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down DivService...")
	grpcServer.GracefulStop()
	log.Println("DivService gracefully stopped.")
}
