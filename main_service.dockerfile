FROM golang:1.22.1-bookworm

RUN go version
ENV GOPATH=/

WORKDIR /soa

COPY ./main_service/ ./main_service/
COPY ./posts/ ./posts/

COPY ./wait-for-postgres.sh ./wait-for-postgres.sh
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o main ./cmd/main.go

CMD ["./main"]
