# 🧠 FizzBuzz API – Go + Gin (DDD, Docker, Redis)

An extensible and testable REST API to dynamically generate FizzBuzz sequences.


---
## 🚀 Getting Started

```bash
git clone git@github.com:ivan-commits/fizzbuzz.git
cd fizzbuzz
make start          # Start the app and Redis via Docker
```

App will be available at: http://localhost:8000

---

## 🌐 Endpoints

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
## 📁 Project Structure

```
.
├── cmd/
│   ├── router/                     # Sets up Gin router with injected dependencies
│   │   └── router.go
│   └── server/                     # App entry point
│       ├── main.go
│       └── main_test.go
├── config/                         # Application configuration constants (e.g., Redis addr)
├── internal/
│   └── fizzbuzz/
│       ├── adapter/
│       │   └── redis/              # Redis implementation of StatsRepository
│       │       ├── redis_stats_repository.go
│       │       └── redis_stats_repository_test.go
│       ├── domain/
│       │   ├── contract/           # Domain contract (ports,interfaces): Generator, Validator, Stats
│       │   │   ├── generator.go
│       │   │   ├── stats.go
│       │   │   └── validator.go
│       │   └── model/              # Domain model (DTOs)
│       │       └── fizzbuzz_dto.go
│       ├── interface/
│       │   └── http/
│       │       ├── handler/        # HTTP handlers
│       │       │   ├── fizzbuzz_handler.go
│       │       │   └── fizzbuzz_handler_test.go
│       │       └── mapper/         # Maps HTTP <-> Domain (requests and responses)
│       │           ├── fizzbuzz_mapper.go
│       │           ├── fizzbuzz_mapper_test.go
│       │           ├── fizzbuzz_response.go
│       │           ├── fizzbuzz_response_test.go
│       │           ├── stats_response.go
│       │           └── stats_response_test.go
│       └── usecase/                # Core business logic implementation
│           ├── default_generator.go
│           ├── default_generator_test.go
│           ├── default_validator.go
│           └── default_validator_test.go
├── redis.conf                     # Redis config (prod tuning)
├── docker-compose.yml             # Docker setup (API + Redis)
├── Dockerfile                     # Multi-stage production build
├── Makefile                       # Dev & build commands
├── go.mod
├── go.sum
└── Readme.md
```
---

## 🐳 Docker

```bash
docker compose up --build
```

Redis listens on `localhost:6379`, the API on `localhost:8000`.

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
make test          # Run all tests
make start         # Start FizzBuzz app and Redis from Docker
make stop          # Stop all containers
make restart       # Stop and clean Redis volume, then rebuild and restart
make kill          # Remove containers, volumes
make logs          # Displays Docker service logs
make redis-cli     # Opens a Redis shell in the container
```
