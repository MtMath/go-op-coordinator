package main

import (
	"os"

	"github.com/joho/godotenv"
	coordpb "notask/op-coordinator/api/coordpb"
	"notask/op-coordinator/internal/coordinator"
	"notask/op-coordinator/internal/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	godotenv.Load()

	addAddr := os.Getenv("ADD_ADDR")
	subAddr := os.Getenv("SUB_ADDR")
	mulAddr := os.Getenv("MUL_ADDR")
	divAddr := os.Getenv("DIV_ADDR")
	coordinatorAddr := os.Getenv("COORDINATOR_ADDR")

	grpcServer := server.New(coordinatorAddr)

	svc := coordinator.NewCoordinatorService(
		addAddr,
		subAddr,
		mulAddr,
		divAddr,
	)

	grpcServer.RegisterService(func(s *grpc.Server) {
		coordpb.RegisterCoordinatorServiceServer(s, svc)
		reflection.Register(s)
	})

	grpcServer.Start("CoordinatorService")
	grpcServer.WaitForShutdown()
}
