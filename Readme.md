# FizzBuzz API – Golang 1.22.2

Une API REST écrite en Go qui expose une version paramétrable du jeu FizzBuzz, avec un endpoint de statistiques.

---

## 🔧 Fonctionnalités

- Endpoint `/fizzbuzz` paramétrable :
  - `int1`, `int2`, `limit`, `str1`, `str2`
  - Exemple : `3`, `5`, `100`, `"fizz"`, `"buzz"`
- Endpoint `/stats` :
  - Affiche la requête la plus fréquente avec son nombre de hits
- Architecture modulaire (handlers, logique métier, stats, modèles)
- Tests unitaires complets
- Makefile pour automatiser les tâches
- Build binaire dans `./bin/`

---

## 🗂️ Structure du projet

```
fizzbuzz-go/
├── cmd/server/          # main.go + router
├── internal/
│   ├── fizzbuzz/        # logique métier
│   ├── handler/         # handlers HTTP
│   └── stats/           # stats en mémoire
├── pkg/model/           # struct FizzBuzzRequest
├── bin/                 # dossier de build (créé par make build)
├── go.mod / go.sum
├── Makefile
└── README.md
```

---

## 🚀 Lancer l'application

### Méthode 1 : exécution directe

```bash
go run ./cmd/server
```

### Méthode 2 : avec Makefile

```bash
make run      # Lance le serveur
```

---

## 📦 Endpoints

### GET `/fizzbuzz`

| Paramètre | Type   | Description                  |
|-----------|--------|------------------------------|
| int1      | int    | multiple remplacé par str1   |
| int2      | int    | multiple remplacé par str2   |
| limit     | int    | borne supérieure             |
| str1      | string | mot à afficher pour int1     |
| str2      | string | mot à afficher pour int2     |

**Exemple** :

```bash
curl "http://localhost:8000/fizzbuzz?int1=3&int2=5&limit=15&str1=fizz&str2=buzz"
```

**Réponse** :

```json
["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz"]
```

---

### GET `/stats`

Renvoie la combinaison la plus utilisée et son nombre de hits.

```bash
curl http://localhost:8000/stats
```

**Réponse** :

```json
{
  "request": {
    "int1": 3,
    "int2": 5,
    "limit": 15,
    "str1": "fizz",
    "str2": "buzz"
  },
  "hits": 7
}
```

---

## 🧪 Tests

Lance tous les tests avec :

```bash
make test
```

Cela couvre :
- La logique FizzBuzz (`internal/fizzbuzz`)
- Le comptage des stats (`internal/stats`)
- Les handlers HTTP (`internal/handler`)
- Le routeur (`cmd/server`)
- La structure du modèle (`pkg/model`)


## 📌 Prérequis

- Go ≥ 1.22.2
- `make` installé (pour utiliser le Makefile, sinon les commandes Go classiques suffisent)

---

## 📄 Licence

Libre d’utilisation, open-source pour usage technique ou démonstration.

