FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/bin/coordinator ./cmd/coordinator_server
RUN go build -o /app/bin/add ./cmd/add_server
RUN go build -o /app/bin/sub ./cmd/sub_server
RUN go build -o /app/bin/mul ./cmd/mul_server
RUN go build -o /app/bin/div ./cmd/div_server


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/ .

EXPOSE 5000
