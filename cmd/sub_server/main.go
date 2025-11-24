package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	subpb "notask/op-coordinator/api/subpb"
	"notask/op-coordinator/internal/server"

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
	godotenv.Load()
	addr := os.Getenv("SUB_ADDR")

	grpcServer := server.New(addr)
	grpcServer.RegisterService(func(s *grpc.Server) {
		subpb.RegisterSubServiceServer(s, &SubServer{})
	})
	grpcServer.Start("SubService")
	grpcServer.WaitForShutdown()
}
