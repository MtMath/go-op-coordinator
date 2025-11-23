package main

import (
	"log"
	"net"
	"os"

	coordpb "notask/op-coordinator/api/coordpb"
	"notask/op-coordinator/internal/coordinator"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	godotenv.Load()

	addAddr := os.Getenv("ADD_ADDR")
	subAddr := os.Getenv("SUB_ADDR")
	mulAddr := os.Getenv("MUL_ADDR")
	divAddr := os.Getenv("DIV_ADDR")

	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Error opening port 5000: %v", err)
	}

	grpcServer := grpc.NewServer()

	svc := coordinator.NewCoordinatorService(
		addAddr,
		subAddr,
		mulAddr,
		divAddr,
	)

	coordpb.RegisterCoordinatorServiceServer(grpcServer, svc)

	reflection.Register(grpcServer)

	log.Println("Coordinator running on :5000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error: %v", err)
	}

	grpcServer.Serve(lis)
}
