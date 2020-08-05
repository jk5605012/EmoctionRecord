FROM golang:1.14

WORKDIR /go/src/gin-test-example

COPY . .

RUN go build .
RUN go test -v ./...
