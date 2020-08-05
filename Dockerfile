FROM golang:1.14

WORKDIR /go/src/gin-test-example

COPY . .

RUN export DOCKERTEST_BIND_LOCALHOST=true
RUN go build .
RUN go test -v ./...