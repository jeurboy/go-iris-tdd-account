FROM golang:1.12
RUN apt-get update && apt-get install -y vim
WORKDIR /usr/src/myapp

ENV GOPATH=/usr
ENV GO111MODULE=on

RUN go get github.com/oxequa/realize