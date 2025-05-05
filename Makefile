install:
	docker compose --profile install up

dev:
	docker compose --profile dev up

cli:
	docker compose --profile cli up -d && docker exec -it cli npx
mysqlcli:
	docker compose exec -it database mariadb -uroot -prootpass api
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
	docker compose exec cli go run apps/api seed
purge:
	docker compose exec cli go run apps/api purge