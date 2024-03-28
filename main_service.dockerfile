FROM golang:1.22.1-bookworm

RUN go version
ENV GOPATH=/

WORKDIR /soa

COPY ./main_service/ ./main_service/
COPY ./posts/ ./posts/

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
COPY ./.env ./.env

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN go mod download
RUN go build -o main ./main_service/cmd/main.go

CMD ["./main"]
