# --- Stage 1 : Build ---
    FROM golang:1.22-alpine AS builder

    WORKDIR /app
    
    # Copy dependency files only
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy source file
    COPY . .
    
    # Compiling the binary
    RUN go build -o fizzbuzz ./cmd/server
    
# --- Stage 2 : Runtime ---
    FROM alpine:3.19
    
    WORKDIR /app
    
    # Copy compiled binary
    COPY --from=builder /app/fizzbuzz .
    
    # Non-root user
    RUN adduser -D -u 10001 fizzbuzz
    USER fizzbuzz
    
    EXPOSE 8000
    
    CMD ["./fizzbuzz"]