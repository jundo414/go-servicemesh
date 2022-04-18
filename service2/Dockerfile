FROM golang:alpine3.15 AS builder

WORKDIR /app

# download the required Go dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64

RUN go build \
    -o /go/bin/main \
    -ldflags '-s -w'

WORKDIR /app/tools
RUN go build \
    -o /go/bin/initialize_redis \
    -ldflags '-s -w'

FROM scratch as runner

COPY --from=builder /go/bin/initialize_redis /app/initialize_redis
COPY --from=builder /go/bin/main /app/main

EXPOSE 19003

RUN ["/app/initialize_redis"]
ENTRYPOINT ["/app/main"]