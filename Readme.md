# 🧠 FizzBuzz API – Architecture DDD avec Go + Gin

Cette application expose une API REST permettant de générer une séquence FizzBuzz configurable.  
Elle est structurée selon une architecture **DDD (Domain-Driven Design) simplifiée**, pour une meilleure lisibilité, testabilité et évolutivité.

---

## 🚀 Endpoints

### `GET /fizzbuzz`

| Paramètre | Type   | Description                                          |
|-----------|--------|------------------------------------------------------|
| int1      | int    | Diviseur pour `str1`                                 |
| int2      | int    | Diviseur pour `str2`                                 |
| limit     | int    | Limite supérieure de la séquence (inclus)           |
| str1      | string | Remplacement pour les multiples de `int1`           |
| str2      | string | Remplacement pour les multiples de `int2`           |

### `GET /stats`

Renvoie la requête `/fizzbuzz` la plus fréquente avec son nombre d'appels.

---

## 🧱 Structure du projet (DDD)

```bash
fizzbuzz/
├── cmd/                          # Point d'entrée (main.go)
│   └── server/
├── internal/
│   └── fizzbuzz/
│       ├── domain/               # Modèle métier pur (structs, interfaces)
│       ├── application/          # Orchestration métier + implémentations
│       └── interfaces/
│           └── http/             # Adaptateurs HTTP (handlers, parsing)
├── go.mod
└── README.md
```

---

## 📁 Rôle de chaque dossier

| Dossier                        | Rôle                                                                 |
|-------------------------------|----------------------------------------------------------------------|
| `cmd/`                         | Contient les exécutables de l’app (ex: API HTTP)                    |
| `domain/`                      | Définit les modèles métier et interfaces (pas de dépendance externe)|
| `application/`                | Contient l’orchestration métier et les implémentations concrètes    |
| `interfaces/http/`            | Gère les requêtes entrantes (HTTP avec Gin)                         |

---

## ✅ Principes appliqués

- **SOLID** : séparation des responsabilités, inversion des dépendances
- **DDD** : chaque couche a un rôle précis, découplé
- **Testabilité** : toutes les couches sont injectables et testables

---

## ▶️ Lancer le projet

```bash
go run ./cmd/server
```

---

Tu veux une version avec `Makefile`, tests, Docker ou Swagger ? Je peux te générer tout ça.
