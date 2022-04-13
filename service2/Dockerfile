# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./


COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY proto/ ./
COPY models/ ./

EXPOSE 19003/tcp

CMD [ "go", "run", "initialize_redis.go" ]
CMD [ "go", "run", "main.go" ]