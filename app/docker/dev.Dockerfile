FROM golang:1.19

WORKDIR /go/src/app
COPY go.mod go.sum ./

RUN go mod download
