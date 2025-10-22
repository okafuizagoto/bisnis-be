# Step 1: build binary
FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/http

# Step 2: run binary
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/main .

# expose port http kamu (misalnya 8080)
EXPOSE 8080  

CMD ["./main"]
