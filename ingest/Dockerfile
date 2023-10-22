FROM golang:1.21-bullseye

WORKDIR /opt/stocks

COPY src src 

RUN go build -C src/app

ENTRYPOINT [ "/bin/bash" ]