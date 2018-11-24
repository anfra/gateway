FROM alpine:latest
MAINTAINER Anton Frank
COPY gateway /usr/bin
EXPOSE 8080
CMD ["/usr/bin/gateway"]