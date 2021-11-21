package main

import (
	"context"
	"log"
	"net"

	pb "go_server/proto/go_server"
	calculator "go_server/service"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedCalculatorServiceServer
}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	return calculator.Add(ctx, in)
}

func (s *server) Multiple(ctx context.Context, in *pb.MultipleRequest) (*pb.MultipleResponse, error) {
	return calculator.Multiple(ctx, in)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
