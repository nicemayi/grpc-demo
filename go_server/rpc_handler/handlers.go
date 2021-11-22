package rpc_handler

import (
	"context"
	pb "go_server/proto/go_server"
	calculator "go_server/service"

	"go_server/service/dto"
	"log"
)

type rpcHandler struct {
	pb.UnimplementedCalculatorServiceServer
}

func NewRpcHandler() *rpcHandler {
	return &rpcHandler{}
}

func (r *rpcHandler) Add(ctx context.Context, in *pb.AddRequest) (*pb.CommonResponse, error) {
	res, err := calculator.Add(ctx, &dto.AddRequest{
		A: int(in.GetA()),
		B: int(in.GetB()),
		C: int(in.GetC()),
	})
	if err != nil {
		log.Fatalf("can not add with error %v", err)
	}
	return &pb.CommonResponse{
		Results: int32(res),
		Message: "Success",
	}, nil
}

func (r *rpcHandler) Multiple(ctx context.Context, in *pb.MultipleRequest) (*pb.CommonResponse, error) {
	res, err := calculator.Multiple(ctx, &dto.MultipleRequest{
		A: int(in.GetA()),
		B: int(in.GetB()),
	})
	if err != nil {
		log.Fatalf("can not multiple with error %v", err)
	}
	return &pb.CommonResponse{
		Results: int32(res),
		Message: "Success",
	}, nil
}

func (r *rpcHandler) Fib(ctx context.Context, in *pb.FibRequest) (*pb.CommonResponse, error) {
	res, err := calculator.Fib(ctx, &dto.FibRequest{
		Number: int(in.GetNumber()),
	})
	if err != nil {
		log.Fatalf("can not fib with error %v", err)
	}
	return &pb.CommonResponse{
		Results: int32(res),
		Message: "Success",
	}, nil
}
