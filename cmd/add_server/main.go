package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"

	addpb "notask/op-coordinator/api/addpb"
	"notask/op-coordinator/internal/server"

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
	godotenv.Load()
	addr := os.Getenv("ADD_ADDR")

	grpcServer := server.New(addr)

	grpcServer.RegisterService(func(s *grpc.Server) {
		addpb.RegisterAddServiceServer(s, &AddServer{})
	})

	grpcServer.Start("AddService")
	grpcServer.WaitForShutdown()
}
