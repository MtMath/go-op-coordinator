package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	divpb "notask/op-coordinator/api/divpb"
	"notask/op-coordinator/internal/server"

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
	godotenv.Load()
	addr := os.Getenv("DIV_ADDR")

	grpcServer := server.New(addr)
	grpcServer.RegisterService(func(s *grpc.Server) {
		divpb.RegisterDivServiceServer(s, &DivServer{})
	})
	grpcServer.Start("DivService")
	grpcServer.WaitForShutdown()
}
