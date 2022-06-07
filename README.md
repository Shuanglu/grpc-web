# Paremeters
- -role "the role to be executed"
- -server_addr "the server address. Only valid when the role is 'server'"
- -version "the server version. Only valid when the role is 'client'"
- -http_port "the http server port"
- -grpc_port "the grpc server port"

# Build
- go build -o \<file name\> .


# Usage
- Start a client
  - \<executable file\> -role client -server_addr <server fqdn> -http_port \<default: 8001> -grpc_port \<default: 50051>
- Start a server
  - \<executable file\> -role server -version v1 -http_port \<default:8001> -grpc_port \<default: 50051>
