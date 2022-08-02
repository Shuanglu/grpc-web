# GRPCWeb
## Parameters
- -role "the role to be executed"
- -client_server_addr "the server address. Only valid when the role is 'client'"
- -version "the server version. Only valid when the role is 'server'"
- -server_http_port "the http port is being listened. Only used for server. Default is 8001"
- -server_grpc_port "the grpc port is being listened. Only used for server.. Default is 50051"
- -client_http_port "the http port of destination. Only used for client. Default is 8001"
- -client_grpc_port "the grpc port of destination. Only used for client. Default is 50051"
- -client_header_host "the host to be added to the header"
- -client_grpc "whether to start grpc client. Default is false"
- -client_http "whether to start http client. Default is true"
- -client_success_request_total "the total requests will be sent to server after the connection is established. Default is 20. 0 is endless loop(This will generate lots of data and send to datadogs. Be careful!)"
- -client_failure_request_total "the total requests will be sent to server with connection failure. Default is 20. 0 is endless loop(This will generate lots of data and send to datadogs. Be careful!)"

## Build
- go build -o \<file name\> .


## Usage
- Start a client
  ```
  <executable file> -role client -server_addr <server fqdn> -client_http_port <default: 8001> -client_grpc_port <default: 50051>
  ```
- Start a server
  ```
  <executable file> -role server -version v1 -server_http_port <default: 8001> -server_grpc_port <default: 50051>
  ```
