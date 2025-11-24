package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	mulpb "notask/op-coordinator/api/mulpb"
	"notask/op-coordinator/internal/server"

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
	godotenv.Load()
	addr := os.Getenv("MUL_ADDR")

	grpcServer := server.New(addr)
	grpcServer.RegisterService(func(s *grpc.Server) {
		mulpb.RegisterMulServiceServer(s, &MulServer{})
	})
	grpcServer.Start("MulService")
	grpcServer.WaitForShutdown()
}
