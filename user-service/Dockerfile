FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/main .

COPY --from=builder /app/main .

COPY --from=builder /app/user_database.sql ./user_database.sql

EXPOSE 8081
CMD ["./main"]