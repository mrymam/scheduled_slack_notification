FROM golang:1.19 as builder
RUN mkdir /build
COPY . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GO111MODULE=on go build -o ./main ./main.go

FROM scratch
COPY --from=builder /build/main /go/bin/main
WORKDIR /app
CMD ["/go/bin/main"]
