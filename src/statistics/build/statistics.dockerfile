FROM golang:1.22.1-bookworm

RUN go version
ENV GOPATH=/

WORKDIR /statistics

COPY ./ ./

RUN go mod download && go build -o main ./cmd/main.go

CMD ["./main"]
