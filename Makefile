# Makefile – FizzBuzz API

run:
	go run ./cmd/server

test:
	go test ./...

up:
	docker compose up --build -d

down:
	docker compose down

kill:
	docker ps -qf "name=fizzbuzz" | xargs -r docker rm -f
	docker ps -qf "ancestor=redis:7-alpine" | xargs -r docker rm -f
	docker compose down -v --rmi all --remove-orphans
	docker system prune -f

logs:
	docker compose logs -f --tail=100

redis-cli:
	docker exec -it $$(docker ps -qf "ancestor=redis:7-alpine") redis-cli