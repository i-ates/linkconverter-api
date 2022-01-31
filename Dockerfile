FROM golang:1.17 AS builder

RUN apt-get update \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go test -v ./...

WORKDIR /app

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o main .

FROM alpine

WORKDIR /app

COPY --from=builder /app/configs/ ./configs/
COPY --from=builder /app/main ./main

EXPOSE 2626

ENTRYPOINT ["/app/main"]
