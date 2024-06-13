FROM golang:1.22.1-bookworm

RUN go version
ENV GOPATH=/

WORKDIR /posts

RUN apt-get update && apt-get -y install postgresql-client

COPY ./ ./

RUN go mod download && go build -o main ./cmd/.

CMD ["./main"]
