FROM debian:stretch-slim
MAINTAINER Anton Frank
ENV LISTENER_PORT=8081
ENV INFLUX_URL=http://influxdb:8086
COPY gateway /usr/bin/
EXPOSE ${LISTENER_PORT}
CMD ["/usr/bin/gateway"]