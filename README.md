# FizzBuzz API (PHP Symfony + Docker)

## Aperçu

Cette application implémente une version paramétrable du fameux FizzBuzz. Elle expose deux endpoints :

Le fizz-buzz original consiste à écrire tous les nombres de 1 à 100, puis à remplacer tous les multiples de 3 par « fizz », tous les multiples de 5 par « buzz » et tous les multiples de 15 par « fizzbuzz ».
Le résultat ressemblerait à ceci : « 1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,… ».

Votre objectif est d'implémenter un serveur web qui exposera un point de terminaison d'API REST qui :
- Accepte cinq paramètres : trois entiers int1, int2 et limit, et deux chaînes str1 et str2.
- Renvoie une liste de chaînes de caractères contenant des nombres de 1 à limit, où : tous les multiples de int1 sont remplacés par str1, tous les multiples de int2 sont remplacés par str2, tous les multiples de int1 et int2 sont remplacés par str1str2.
Le serveur doit être :
- Prêt pour la production
- Facile à maintenir par d'autres développeurs
Bonus : ajouter un point de terminaison de statistiques permettant aux utilisateurs de connaître les requêtes les plus fréquentes. Ce point de terminaison doit :
- N'accepter aucun paramètre
- Renvoyer les paramètres correspondant à la requête la plus utilisée, ainsi que le nombre de hits pour cette requête.

## Installation et Exécution (via Docker uniquement)

1. **Cloner le dépôt** :

   ```bash
   git clone https://votre-repo/fizzbuzz.git
   cd fizzbuzz
   ```

2. **Construire l’image Docker et lancer le service** :

   ```bash
   docker-compose build
   docker-compose up -d
   ```

3. **Accéder à l’application** :

   - L’API est exposée sur le port 8000 (modifiable dans `docker-compose.yml`).
   - Rendez-vous sur [http://localhost:8000](http://localhost:8000) pour accéder au projet.

## Utilisation et Endpoints

### 1. Endpoint : `/fizzbuzz`
- **Méthode** : `GET`  
- **Paramètres (query params)** :  
  - `int1` (ex. `3`)  
  - `int2` (ex. `5`)  
  - `limit` (ex. `15`)  
  - `str1` (ex. `fizz`)  
  - `str2` (ex. `buzz`)

**Fonctionnement** :  
Retourne un tableau JSON de valeurs allant de `1` à `limit`, en appliquant la logique FizzBuzz paramétrée. Les multiples de `int1` et `int2` sont remplacés par la concaténation `str1str2`.

#### Exemple d’appel :

```bash
GET http://localhost:8000/fizzbuzz?int1=3&int2=5&limit=15&str1=fizz&str2=buzz
```

**Réponse** (extrait) :

```json
[
  "1",
  "2",
  "fizz",
  "4",
  "buzz",
  "fizz",
  "7",
  "8",
  "fizz",
  "buzz",
  "11",
  "fizz",
  "13",
  "14",
  "fizzbuzz"
]
```

### 2. Endpoint : `/stats`
- **Méthode** : `GET`  
- **Paramètres** : *Aucun*

**Fonctionnement** :  
Retourne un objet JSON correspondant à la combinaison de paramètres FizzBuzz la plus souvent appelée, ainsi que le nombre total d’appels pour cette combinaison.

#### Exemple d’appel :

```bash
GET http://localhost:8000/stats
```

**Réponse** (exemple) :

```json
{
  "int1": 3,
  "int2": 5,
  "limit": 15,
  "str1": "fizz",
  "str2": "buzz",
  "hits": 27
}
```

## Tests

- Les tests (unitaires et/ou fonctionnels) se trouvent généralement dans le dossier `tests/`.
- Pour les exécuter, vous pouvez configurer une commande Docker pour lancer `bin/phpunit` au sein du conteneur, ou lancer `php bin/phpunit` en local si vous avez PHP et les dépendances installées.
