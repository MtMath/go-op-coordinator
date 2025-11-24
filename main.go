package main

import (
	"log"
	"net"
	"os"

	addpb "notask/op-coordinator/api/addpb"
	coordpb "notask/op-coordinator/api/coordpb"
	divpb "notask/op-coordinator/api/divpb"
	mulpb "notask/op-coordinator/api/mulpb"
	subpb "notask/op-coordinator/api/subpb"

	addServer "notask/op-coordinator/cmd/add_server"
	subServer "notask/op-coordinator/cmd/sub_server"
	mulServer "notask/op-coordinator/cmd/mul_server"
	divServer "notask/op-coordinator/cmd/div_server"

	"notask/op-coordinator/internal/coordinator"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func startGRPCServer(name, addr string, register func(*grpc.Server)) {
	go func() {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("Erro iniciando %s na porta %s: %v", name, addr, err)
		}

		grpcServer := grpc.NewServer()
		register(grpcServer)

		reflection.Register(grpcServer)

		log.Printf("%s rodando em %s\n", name, addr)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Erro no %s: %v", name, err)
		}
	}()
}

func main() {

	godotenv.Load()

	addAddr := getenv("ADD_ADDR", ":5001")
	subAddr := getenv("SUB_ADDR", ":5002")
	mulAddr := getenv("MUL_ADDR", ":5003")
	divAddr := getenv("DIV_ADDR", ":5004")
	coordAddr := getenv("COORD_ADDR", ":5000")

	startGRPCServer("AddService", addAddr, func(s *grpc.Server) {
		addpb.RegisterAddServiceServer(s, &addServer.AddServer{})
	})

	startGRPCServer("SubService", subAddr, func(s *grpc.Server) {
		subpb.RegisterSubServiceServer(s, &subServer.SubServer{})
	})

	startGRPCServer("MulService", mulAddr, func(s *grpc.Server) {
		mulpb.RegisterMulServiceServer(s, &mulServer.MulServer{})
	})

	startGRPCServer("DivService", divAddr, func(s *grpc.Server) {
		divpb.RegisterDivServiceServer(s, &divServer.DivServer{})
	})

	startGRPCServer("CoordinatorService", coordAddr, func(s *grpc.Server) {
		svc := coordinator.NewCoordinatorService(addAddr, subAddr, mulAddr, divAddr)
		coordpb.RegisterCoordinatorServiceServer(s, svc)
	})

	log.Println("All microservices are running! Press CTRL+C to exit.")

	select {}
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
