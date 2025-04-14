# Makefile – FizzBuzz API

run:
	go run ./cmd/server

test:
	go test ./...

up:
	docker compose up --build -d

down:
	docker compose down

restart:
	docker compose down --volumes --remove-orphans
	docker volume rm fizzbuzz_redis-data || true
	docker compose up --build -d

kill:
	docker ps -qf "name=fizzbuzz" | xargs -r docker rm -f
	docker ps -qf "ancestor=redis:7-alpine" | xargs -r docker rm -f
	docker compose down --volumes --remove-orphans --rmi all

logs:
	docker compose logs -f --tail=100

redis-cli:
	docker exec -it $$(docker ps -qf "ancestor=redis:7-alpine") redis-cli