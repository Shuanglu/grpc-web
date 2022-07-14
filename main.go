package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"sync"

	"github.com/shuanglu/grpc-web/grpcclient"
	"github.com/shuanglu/grpc-web/grpcserver"
	"github.com/shuanglu/grpc-web/webclient"
	"github.com/shuanglu/grpc-web/webserver"
)

var (
	role               = flag.String("role", "server", "the role to be executed. Options: client/server/both")
	server_http_port   = flag.Int("server_http_port", 8001, "the http server port")
	server_grpc_port   = flag.Int("server_grpc_port", 50051, "the grpc server port")
	version            = flag.String("version", "", "the server version")
	client_server_addr = flag.String("client_server_addr", "localhost", "the server address of destination")
	client_http_port   = flag.Int("client_http_port", 8001, "the http port of destination")
	client_grpc_port   = flag.Int("client_grpc_port", 50051, "the grpc port of destination")
	client_header_host = flag.String("client_header_host", "", "the host to be added to the headr")
	client_http        = flag.Bool("client_http", false, "whether to start http client")
	client_grpc        = flag.Bool("client_grpc", true, "whether to start grpc client")
)

func main() {
	flag.Parse()
	mesh := os.Getenv("mesh")
	ip := os.Getenv("POD_IP")
	log.Printf("This is an app running in the %q mesh", mesh)
	var wg sync.WaitGroup
	wg.Add(1)
	go grpc_run(fmt.Sprintf("%s:%d", *client_server_addr, *client_grpc_port), *client_header_host, mesh, ip, *role)
	wg.Add(1)
	go http_run(fmt.Sprintf("http://%s:%d", *client_server_addr, *client_http_port), *client_header_host, mesh, ip, *role)
	wg.Wait()
}

func grpc_run(dest string, host string, mesh string, ip string, role string) {
	if role == "server" {
		grpcserver.Run(server_grpc_port, version, mesh, ip)
	} else if role == "client" {
		if *client_grpc {
			grpcclient.Run(dest, host, mesh, ip)
		}
	} else {
		go grpcserver.Run(server_grpc_port, version, mesh, ip)
		if *client_grpc {
			grpcclient.Run(dest, host, mesh, ip)
		}
	}
}

func http_run(dest string, host string, mesh string, ip string, role string) {

	if role == "server" {
		webserver.Run(server_http_port, version, mesh, ip)
	} else if role == "client" {
		if *client_http {
			go webclient.Run(dest, host, mesh, ip)
		}
	} else {
		go webserver.Run(server_http_port, version, mesh, ip)
		if *client_http {
			go webclient.Run(dest, host, mesh, ip)
		}
	}
}
