.PHONY: up stop down destroy

up:
	docker compose -f compose.yaml up -d

stop:
	docker compose -f compose.yaml stop