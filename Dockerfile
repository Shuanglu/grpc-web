FROM ubuntu
COPY output/grpcweb /bin/
RUN chmod 755 /bin/grpcweb 
# below is for alpine
# && mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 
ENTRYPOINT [ "grpcweb" ] 