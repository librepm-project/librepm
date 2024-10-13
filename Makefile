install:
	docker compose --profile install up

dev:
	docker compose --profile dev up

cli:
	docker compose --profile cli up -d && docker exec -it cli npx