# Paremeters
- -role "the role to be executed"
- -server_addr "the server address. Only valid when the role is 'server'"
- -version "the server version. Only valid when the role is 'client'"
- -server_http_port "the http server port"
- -server_grpc_port "the grpc server port"
- -client_http_port "the http port of destination"
- -client_grpc_port "the grpc port of destination"

# Build
- go build -o \<file name\> .


# Usage
- Start a client
  - \<executable file\> -role client -server_addr <server fqdn> -client_http_port \<default: 8001> -client_grpc_port \<default: 50051>
- Start a server
  - \<executable file\> -role server -version v1 -server_http_port \<default:8001> -server_grpc_port \<default: 50051>
