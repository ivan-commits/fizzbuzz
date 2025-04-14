# 🧠 FizzBuzz API – Go + Gin (DDD, Docker, Redis)

An extensible and testable REST API to dynamically generate FizzBuzz sequences.

---

## 🚀 Endpoints

- `GET /fizzbuzz`  
  ➤ Generates the FizzBuzz sequence according to your settings:

  | Parameter | Type   | Description                         |
  |-----------|--------|-------------------------------------|
  | int1      | int    | Divider for `str1`                |
  | int2      | int    | Divider for `str2`                |
  | limit     | int    | Upper limit (inclusive)         |
  | str1      | string | Replaces multiples of `int1`    |
  | str2      | string | Replaces multiples of `int2`    |

  
  ➤ Response :
```json
{
  "response": ["1", "2", "fizz", "4", "buzz", "fizz", ...]
}
```

- `GET /stats`  
  ➤ Returns the most frequent FizzBuzz query + number of calls.
  
  ➤ Response :
```json
{
  "key": "int1=3&int2=5&limit=100&str1=fizz&str2=buzz",
  "count": 42
}
```
---

## 🧱 Structure DDD

```
internal/
├── fizzbuzz/
│   ├── domain/        # Business interfaces (contract) & DTOs
│   ├── usecase/       # Business logic implemented
│   ├── adapter/redis/ # Redis Persistence (Stats)
│   └── interface/
│       └── http/
│           ├── handler/ # Handlers REST
│           └── mapper/  # Mapping HTTP <-> Domain
```

---

## 🐳 Docker

```bash
docker compose up --build
```

Redis listens on `localhost:6379`, the API on `localhost:8000`.

---

## ⚙️ Production

- `vm.overcommit_memory=1` recommandé pour Redis
- `GIN_MODE=release` (already injected in `docker-compose.yml`)

---

## ✅ Tests

```bash
go test ./...
```

All unit and integration tests are covered (handlers, usecase, Redis, mapping...).

---

## 📦 Configuration

Values ​​centralized in `config/config.go` :
```go
const (
	DefaultPort          = ":8000"
	DefaultRedisAddr     = "redis:6379"
	LocalHostRedisAddr   = "localhost:6379"
	DefaultRedisDB       = 1
	AlternateRedisTestDB = 2
)

```

## 🛠️ Makefile

Useful commands :

```bash
make run           # Run the app locally
make test          # Run all tests
make up            # Start FizzBuzz app and Redis from Docker
make down          # Stop all containers
make restart       # Stop and clean Redis volume, then rebuild and restart
make kill          # Remove containers, volumes
make logs          # Displays Docker service logs
make redis-cli     # Opens a Redis shell in the container
```