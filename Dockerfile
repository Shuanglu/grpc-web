FROM alpine
COPY ./grpcweb /bin/
RUN chmod 755 /bin/grpcweb && ls /bin/grpcweb
ENTRYPOINT [ "tail", "-f", "/dev/null" ] 