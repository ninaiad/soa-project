FROM golang:1.22.1-bookworm

RUN go version
ENV GOPATH=/

WORKDIR /soa

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN go mod download
RUN go build -o main ./cmd/main.go

CMD ["./main"]

