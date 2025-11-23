package main

import (
	"context"
	"log"
	"net"

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

	log.Println("SubService running on :5002")
	grpcServer.Serve(lis)
}
