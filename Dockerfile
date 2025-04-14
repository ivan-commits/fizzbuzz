# --- Stage 1 : Build ---
    FROM golang:1.22-alpine AS builder

    WORKDIR /app
    
    # Copie des fichiers de dépendances uniquement
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copie du code source
    COPY . .
    
    # Compilation du binaire
    RUN go build -o fizzbuzz ./cmd/server
    
# --- Stage 2 : Runtime ---
    FROM alpine:3.19
    
    WORKDIR /app
    
    # Copie du binaire compilé
    COPY --from=builder /app/fizzbuzz .
    
    # Utilisateur non-root (optionnel mais recommandé)
    RUN adduser -D -u 10001 fizzbuzz
    USER fizzbuzz
    
    EXPOSE 8000
    
    CMD ["./fizzbuzz"]