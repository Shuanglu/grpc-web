package main

import (
	"flag"
	"fmt"

	"sync"
	"time"

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
	var wg sync.WaitGroup
	if *role == "server" {
		wg.Add(1)
		go grpcserver.Run(server_grpc_port, version)
		wg.Add(1)
		go webserver.Run(server_http_port, version)
		wg.Wait()
	} else if *role == "client" {
		for {
			if *client_grpc {
				go grpcclient.Run(fmt.Sprintf("%s:%d", *client_server_addr, *client_grpc_port), *client_header_host)
			}
			if *client_http {
				go webclient.Run(fmt.Sprintf("http://%s:%d", *client_server_addr, *client_http_port), *client_header_host)
			}
			time.Sleep(5 * time.Second)
		}
	} else if *role == "both" {
		wg.Add(1)
		go grpcserver.Run(server_grpc_port, version)
		wg.Add(1)
		go webserver.Run(server_http_port, version)

		for {
			if *client_grpc {
				go grpcclient.Run(fmt.Sprintf("%s:%d", *client_server_addr, *client_grpc_port), *client_header_host)
			}
			if *client_http {
				go webclient.Run(fmt.Sprintf("http://%s:%d", *client_server_addr, *client_http_port), *client_header_host)
			}
			time.Sleep(5 * time.Second)
		}
		wg.Wait()
	}

}
