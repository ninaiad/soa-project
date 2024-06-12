FROM golang:1.22.1-bookworm

RUN go version
ENV GOPATH=/

WORKDIR /test

COPY ./ ./

RUN go mod download

CMD ["go", "test", "-v", "./..."]