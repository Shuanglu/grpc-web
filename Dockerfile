FROM alpine
COPY ./grpcweb .
CMD ['./grpcweb']
