package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	subpb "notask/op-coordinator/api/subpb"

	"google.golang.org/grpc"
)

type SubServer struct {
	subpb.UnimplementedSubServiceServer
}

func (s *SubServer) Compute(ctx context.Context, req *subpb.OperationRequest) (*subpb.OperationResponse, error) {
	return &subpb.OperationResponse{
		Result: req.A - req.B,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5002")
	if err != nil {
		log.Fatalf("Error opening port 5002: %v", err)
	}

	grpcServer := grpc.NewServer()
	subpb.RegisterSubServiceServer(grpcServer, &SubServer{})

	go func() {
		log.Println("SubService running on :5002")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down SubService...")
	grpcServer.GracefulStop()
	log.Println("SubService gracefully stopped.")
}
