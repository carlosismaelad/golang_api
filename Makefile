.PHONY: up stop build

up:
	docker compose -f compose.yaml up -d
stop:
	docker compose -f compose.yaml stop
build:
	docker build -t go-crud-api .