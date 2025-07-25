FROM golang:1.24.4 AS builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Include the Firebase key
COPY firebase-service-account.json ./firebase-service-account.json

# Build from cmd/
RUN go build -o main ./cmd

FROM debian:bookworm-slim
WORKDIR /root/

# Copy binary and config file
COPY --from=builder /app/main .
COPY --from=builder /app/firebase-service-account.json .

# Set .env if needed manually OR mount it in docker-compose
EXPOSE 8080
CMD ["./main"]
