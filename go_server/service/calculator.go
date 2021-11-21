package service

import (
	"context"
	"log"

	pb "go_server/proto/go_server"
)

func Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("In go_server/service/calculator.go, Received: %v, %v, %v", in.GetA(), in.GetB(), in.GetC())
	return &pb.AddResponse{
		Results: in.GetA() + in.GetB() + in.GetC(),
		Message: "Succuess",
	}, nil
}

func Multiple(ctx context.Context, in *pb.MultipleRequest) (*pb.MultipleResponse, error) {
	log.Printf("In go_server/service/calculator.go, Received: %v, %v, %v", in.GetA(), in.GetB())
	return &pb.MultipleResponse{
		Results: in.GetA() * in.GetB(),
		Message: "Succuess",
	}, nil
}
