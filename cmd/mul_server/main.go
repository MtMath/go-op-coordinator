package main

import (
	"context"
	"log"
	"net"

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

	log.Println("MulService running on :5003")
	grpcServer.Serve(lis)
}
