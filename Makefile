install:
	docker compose --profile install up

dev:
	docker compose --profile dev up

cli:
	docker compose --profile cli up -d && docker exec -it librepm_cli npx
mysqlcli:
	docker compose exec -it database librepm_mariadb -uroot -prootpass api
open:
	@URL="http://localhost:8081"; \
	if [[ "$$OSTYPE" == "linux-gnu"* ]]; then \
		xdg-open "$$URL"; \
	elif [[ "$$OSTYPE" == "darwin"* ]]; then \
		open "$$URL"; \
	else \
		echo "Unsupported OS. Please open the URL manually: $$URL"; \
	fi
seed:
	docker compose --profile cli exec librepm_cli go run apps/api seed /app/apps/api/seed.yaml
purge:
	docker compose exec librepm_cli go run apps/api purge