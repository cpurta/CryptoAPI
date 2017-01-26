FROM ubuntu:latest

MAINTAINER Chris Purta cpurta@gmail.com

RUN apt-get update && \
    apt-get -y install ca-certificates && \
    mkdir -p /app

ADD crypto-api /app

RUN chmod +x /app/crypto-api

EXPOSE 443

ENTRYPOINT ["/app/crypto-api"]
