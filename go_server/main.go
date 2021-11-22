package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"go_server/handler"
	pb "go_server/proto/go_server"
	calculator "go_server/service"
	"go_server/service/dto"

	"google.golang.org/grpc"
)

const (
	rpcPort  = ":50051"
	httpPort = ":8080"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedCalculatorServiceServer
}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.CommonResponse, error) {
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

func (s *server) Multiple(ctx context.Context, in *pb.MultipleRequest) (*pb.CommonResponse, error) {
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

func (s *server) Fib(ctx context.Context, in *pb.FibRequest) (*pb.CommonResponse, error) {
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

func main() {
	var wg sync.WaitGroup

	// serve rpc
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		rpcServer := grpc.NewServer()
		pb.RegisterCalculatorServiceServer(rpcServer, &server{})
		lis, err := net.Listen("tcp", rpcPort)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		if err := rpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		log.Printf("server listening at %v", lis.Addr())
	}()

	// server http
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		http.HandleFunc("/add", handler.Add)
		http.HandleFunc("/multiple", handler.Multiple)
		http.HandleFunc("/fib", handler.Fib)
		log.Fatal(http.ListenAndServe(httpPort, nil))
	}()

	wg.Wait()
}
