FROM golang:1.14-alpine

RUN apk add --no-cache go
ADD main.go go.mod go.sum /work/
RUN cd /work && go mod download
RUN cd /work && go build -o /usr/bin/app
ENTRYPOINT /usr/bin/app
