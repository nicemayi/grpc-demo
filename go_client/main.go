/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"sync"
	"time"

	pb "go_client/proto/go_server"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func worker(ctx context.Context, c pb.CalculatorServiceClient, i int) int32 {
	r, err := c.Fib(ctx, &pb.FibRequest{Number: int32(i)})
	if err != nil {
		log.Fatalf("could not fib: %v", err)
	}
	// log.Printf("%d-th result: %v", i, r.GetResults())
	return r.GetResults()
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	r, err := c.Add(ctx, &pb.AddRequest{A: 1, B: 2, C: 3})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Result: %v", r.GetResults())

	startTime := time.Now()
	for i := 0; i < 50; i++ {
		worker(ctx, c, i+1)
	}
	log.Printf("Sync: total time cost: %v seconds", time.Since(startTime).Seconds())

	startTime = time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			worker(ctx, c, i+1)
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Printf("Async: total time cost: %v seconds", time.Since(startTime).Seconds())
}
