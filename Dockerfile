FROM alpine:latest
MAINTAINER Anton Frank
COPY gateway /usr/bin
EXPOSE 8081
CMD ["/usr/bin/gateway"]