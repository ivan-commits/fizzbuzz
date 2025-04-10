APP_NAME=fizzbuzz
MAIN_PKG=./cmd/server
PORT=8000

## 🧪 Lancer tous les tests
test:
	go test ./... -v

## 🚀 Lancer le serveur
run:
	go run $(MAIN_PKG)

## 🆘 Aide
help:
	@echo "Commandes disponibles:"
	@echo "  make run     - Lancer le serveur (port $(PORT))"
	@echo "  make test    - Lancer tous les tests"
