FROM debian:latest

RUN mkdir -p /app
WORKDIR /app

COPY consignment.json /app/consignment.json

ADD consignment.json /app/consignment.json
ADD consignment-cli /app/consignment-cli

ENTRYPOINT ["./consignment-cli"]