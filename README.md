# Paremeters
- -role "the role to be executed"
- -server_addr "the server address. Only valid when the role is 'server'"
- -version "the server version. Only valid when the role is 'server'"
- -http_port "the http server port. Only valid when the role is 'server'"
- -grpc_port "the grpc server port. Only valid when the role is 'server'"

# Build
- go build -o \<file name\> .


# Usage
- Start a client
  - \<executable file\> -role client
- Start a server
  - \<executable file\> -role server -version v1 -http_port 8001 -grpc_port 50051
