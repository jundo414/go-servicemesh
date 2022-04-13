# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./


COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY proto/ ./

EXPOSE 8080/tcp

CMD [ "go", "run", "main.go" ]