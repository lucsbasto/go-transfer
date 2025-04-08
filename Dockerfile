FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o go-transfer ./cmd

FROM alpine:latest

ENV GIN_MODE=release

WORKDIR /root/

COPY --from=builder /app/go-transfer .

COPY .env .

EXPOSE 8080

CMD ["./go-transfer"]
