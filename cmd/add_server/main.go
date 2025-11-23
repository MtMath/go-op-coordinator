package main

import (
	"context"
	"log"
	"net"

	addpb "notask/op-coordinator/api/addpb"

	"google.golang.org/grpc"
)

type AddServer struct {
	addpb.UnimplementedAddServiceServer
}

func (s *AddServer) Compute(ctx context.Context, req *addpb.OperationRequest) (*addpb.OperationResponse, error) {
	return &addpb.OperationResponse{
		Result: req.A + req.B,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Error opening port 5001: %v", err)
	}

	grpcServer := grpc.NewServer()
	addpb.RegisterAddServiceServer(grpcServer, &AddServer{})

	log.Println("AddService running on :5001")
	grpcServer.Serve(lis)
}
