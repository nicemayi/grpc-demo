package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"

	httpHandler "go_server/http_handler"
	rpcHandler "go_server/rpc_handler"
	pb "go_server/proto/go_server"

	"google.golang.org/grpc"
)

func init() {
	if os.Getenv("REDIS_URL") == "" {
		log.Panic("Need REDIS_URL")
	}
	if os.Getenv("REDIS_PORT") == "" {
		log.Panic("Need REDIS_PORT")
	}
	if os.Getenv("RPC_PORT") == "" {
		log.Panic("Need RPC_PORT")
	}
	if os.Getenv("HTTP_PORT") == "" {
		log.Panic("Need HTTP_PORT")
	}
	log.Print("Initial check passed!")
}

var (
	rpcPort  = fmt.Sprintf(":%s", os.Getenv("RPC_PORT"))
	httpPort = fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))
)

func main() {
	var wg sync.WaitGroup

	// serve rpc
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		rpcServer := grpc.NewServer()
		rpcHandeler := rpcHandler.NewRpcHandler()
		pb.RegisterCalculatorServiceServer(rpcServer, rpcHandeler)

		lis, err := net.Listen("tcp", rpcPort)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		if err := rpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		log.Printf("rpc server listening at %v", lis.Addr())
	}()

	// server http
	h := httpHandler.NewHttpHanlder()
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		http.HandleFunc("/add", h.Add)
		http.HandleFunc("/multiple", h.Multiple)
		http.HandleFunc("/fib", h.Fib)
		log.Fatal(http.ListenAndServe(httpPort, nil))
		log.Printf("http server listening at %v", httpPort)
	}()

	wg.Wait()
}
