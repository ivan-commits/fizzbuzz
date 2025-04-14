# 🧠 FizzBuzz API – Go + Gin (DDD, Docker, Redis)

Une API REST extensible et testable pour générer dynamiquement des séquences FizzBuzz.

---

## 🚀 Endpoints

- `GET /fizzbuzz`  
  ➤ Génère la séquence FizzBuzz selon vos paramètres :

  | Paramètre | Type   | Description                         |
  |-----------|--------|-------------------------------------|
  | int1      | int    | Diviseur pour `str1`                |
  | int2      | int    | Diviseur pour `str2`                |
  | limit     | int    | Limite supérieure (incluse)         |
  | str1      | string | Remplace les multiples de `int1`    |
  | str2      | string | Remplace les multiples de `int2`    |

  
  ➤ Réponse :
```json
{
  "response": ["1", "2", "fizz", "4", "buzz", "fizz", ...]
}
```

- `GET /stats`  
  ➤ Retourne la requête FizzBuzz la plus fréquente + nombre d'appels.
  
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
│   ├── domain/        # Interfaces métier (contract) & DTOs
│   ├── usecase/       # Logique métier implémentée
│   ├── adapter/redis/ # Persistance Redis (Stats)
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

Redis écoute sur `localhost:6379`, l’API sur `localhost:8000`.

---

## ⚙️ Production

- `vm.overcommit_memory=1` recommandé pour Redis
- `GIN_MODE=release` (déjà injecté via `docker-compose.yml`)
- Logging prévu pour toutes les erreurs critiques

---

## ✅ Tests

```bash
go test ./...
```

Tous les tests unitaires et intégration sont couverts (handlers, usecase, Redis, mapping...).

---

## 📦 Configuration

Valeurs centralisées dans `config/config.go` :
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

Commandes utiles :

```bash
make run           # Lance l'app localement (Go)
make test          # Lance tous les tests
make up            # Démarre l'app et Redis via Docker
make down          # Stoppe les services
make kill          # Supprime containers, volumes, images et orphelins
make logs          # Affiche les logs des services Docker
make redis-cli     # Ouvre un shell Redis dans le container
```