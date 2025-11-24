package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	mulpb "notask/op-coordinator/api/mulpb"

	"google.golang.org/grpc"
)

type MulServer struct {
	mulpb.UnimplementedMulServiceServer
}

func (s *MulServer) Compute(ctx context.Context, req *mulpb.OperationRequest) (*mulpb.OperationResponse, error) {
	return &mulpb.OperationResponse{
		Result: req.A * req.B,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5003")
	if err != nil {
		log.Fatalf("Error opening port 5003: %v", err)
	}

	grpcServer := grpc.NewServer()
	mulpb.RegisterMulServiceServer(grpcServer, &MulServer{})

	go func() {
		log.Println("MulService running on :5003")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down MulService...")
	grpcServer.GracefulStop()
	log.Println("MulService gracefully stopped.")
}
