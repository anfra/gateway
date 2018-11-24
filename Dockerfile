FROM debian:stretch-slim
MAINTAINER Anton Frank
ENV LISTENER_PORT=8081
COPY gateway /usr/bin/
EXPOSE ${LISTENER_PORT}
CMD ["/usr/bin/gateway"]