FROM golang:alpine3.15 AS builder

WORKDIR /app

# download the required Go dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN ls

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
    -o /go/bin/main \
    -ldflags '-s -w'


FROM scratch as runner

COPY --from=builder /go/bin/main /app/main

EXPOSE 8080

ENTRYPOINT ["/app/main"]