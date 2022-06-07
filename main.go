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
	role        = flag.String("role", "server", "the role to be executed")
	http_port   = flag.Int("http_port", 8001, "the http server port")
	grpc_port   = flag.Int("grpc_port", 50051, "the grpc server port")
	version     = flag.String("version", "", "the server version")
	server_addr = flag.String("server_addr", "", "the server address")
)

func main() {
	flag.Parse()
	var wg sync.WaitGroup
	if *role == "server" {
		wg.Add(1)
		go grpcserver.Run(grpc_port, version)
		wg.Add(1)
		go webserver.Run(http_port, version)
		wg.Wait()
	} else if *role == "client" {
		for {

			go grpcclient.Run(fmt.Sprintf("%s:%d", *server_addr, *grpc_port))

			go webclient.Run(fmt.Sprintf("http://%s:%d", *server_addr, *http_port))

			time.Sleep(5 * time.Second)
		}
	} else if *role == "both" {

		go grpcserver.Run(grpc_port, version)

		go webserver.Run(http_port, version)

		for {
			go grpcclient.Run(fmt.Sprintf("%s:%d", *server_addr, *grpc_port))
			go webclient.Run(fmt.Sprintf("http://%s:%d", *server_addr, *http_port))
			time.Sleep(5 * time.Second)
		}

	}

}
