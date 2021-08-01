FROM alpine:latest
WORKDIR /opt
COPY bin/binary /opt/binary
COPY service/api/swaggerui /swaggerui

CMD ["./binary"]
