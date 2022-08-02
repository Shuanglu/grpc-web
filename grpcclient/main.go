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
package grpcclient

import (
	"context"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/metadata"
)

const (
	name = "world"
)

func Run(grpcAddr string, host string, mesh string, ip string, client_success_request_total int, client_failure_request_total int) error {

	// Set up a connection to the server.
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithAuthority(host), grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(grpcAddr, opts...)
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	successCount := 0
	failureCount := 0
	var wg sync.WaitGroup
	for {
		wg.Add(1)
		go func(successCount *int, failureCount *int) {
			c := pb.NewGreeterClient(conn)

			// Contact the server and print out its response.
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)

			ctx = metadata.AppendToOutgoingContext(ctx, "x-real-ip", ip)

			defer cancel()
			r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
			if err != nil {
				log.Printf("could not greet: %v", err)
				conn, err = grpc.Dial(grpcAddr, opts...)
				if err != nil {
					log.Printf("did not connect: %v", err)
				}
				*failureCount++
			} else {
				log.Printf("GRPC | client is running in the mesh: %q | %s ", mesh, r.GetMessage())
				*successCount++
			}
			wg.Done()
		}(&successCount, &failureCount)
		if client_failure_request_total == 0 || client_success_request_total == 0 {
			time.Sleep(5 * time.Second)
			continue
		} else if successCount == client_success_request_total {
			log.Printf("Sent %d requests to server. Will sleep forever", client_success_request_total)
			time.Sleep(time.Duration(1<<63 - 1))
		} else if failureCount == client_failure_request_total {
			log.Printf("Failed to send %d requests to server. Will sleep forever", client_failure_request_total)
			time.Sleep(time.Duration(1<<63 - 1))
		} else {
			time.Sleep(5 * time.Second)
			continue
		}

	}
	wg.Wait()

	return nil
}
