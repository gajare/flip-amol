# ---------- Build stage ----------
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

# Install git (required for go mod)
RUN apk add --no-cache git

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main

# ---------- Runtime stage ----------
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary
COPY --from=builder /app/main .

# Copy SQL file (FIXED PATH)
COPY --from=builder /app/database.sql ./database.sql

EXPOSE 8080

CMD ["./main"]
