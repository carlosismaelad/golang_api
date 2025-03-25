.PHONY: up stop build

up:
	docker compose -f compose.yaml up -d
stop:
	docker compose -f compose.yaml stop
down:
	docker compose -f compose.yaml down
destroy:
	docker system prune -a
volumes-down:
	docker compose down -v
build-app-image:
	docker build -t go-crud-api .