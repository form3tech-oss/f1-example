FROM golang:1.17
COPY ./ /f1-example
WORKDIR /f1-example
CMD go run ./cmd/service/main.go
