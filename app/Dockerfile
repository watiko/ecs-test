FROM golang:1.14-alpine

RUN apk add --no-cache go

WORKDIR /work

ADD go.mod go.sum /work/
RUN go mod download

ADD main.go /work/
RUN go build -o /usr/bin/app

ENTRYPOINT /usr/bin/app
