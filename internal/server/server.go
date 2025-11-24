package server

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	server *grpc.Server
	lis    net.Listener
}

func New(addr string) *GrpcServer {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error opening port %s: %v", addr, err)
	}

	grpcServer := grpc.NewServer()

	return &GrpcServer{
		server: grpcServer,
		lis:    lis,
	}
}

func (s *GrpcServer) RegisterService(register func(*grpc.Server)) {
	register(s.server)
}

func (s *GrpcServer) Start(serviceName string) {
	go func() {
		log.Printf("%s running on %s", serviceName, s.lis.Addr().String())
		if err := s.server.Serve(s.lis); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}()
}

func (s *GrpcServer) WaitForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	s.server.GracefulStop()
	log.Println("Server gracefully stopped.")
}
