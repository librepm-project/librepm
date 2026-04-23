VERSION := $(shell node -p "require('./package.json').version")

build:
	docker build -f Dockerfile.production \
		-t ghcr.io/librepm-project/librepm:latest \
		-t ghcr.io/librepm-project/librepm:$(VERSION) \
		.

push:
	docker push ghcr.io/librepm-project/librepm:latest
	docker push ghcr.io/librepm-project/librepm:$(VERSION)

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

test:
	docker compose --profile cli exec cli go test apps/api/test/...

seed:
	docker compose --profile cli exec cli go run apps/api seed /app/apps/api/seed.yaml

purge:
	docker compose --profile cli exec cli go run apps/api purge

set_version:
ifndef VERSION
	@echo "Usage: make set_version VERSION=x.y.z"
	@exit 1
endif
	@echo "$(VERSION)" | grep -Eq '^[0-9]+\.[0-9]+\.[0-9]+$$' || { echo "Error: VERSION must be semver (x.y.z), got '$(VERSION)'"; exit 1; }
	npm pkg set version=$(VERSION)
	sed -i '' 's/^\([[:space:]]*version:[[:space:]]*\).*/\1$(VERSION)/' apps/api/app/http/openapi.yaml
	@echo "Version set to $(VERSION)"

storybook:
	docker compose --profile cli run --rm --entrypoint sh cli -c "cd /app && npx nx run frontend:storybook"

e2e:
	docker compose --profile dev up -d
	docker compose --profile e2e run --rm playwright

upgrade:
	docker compose down
	git pull
	docker compose up --build -d

login:
ifndef EMAIL
	@echo "Usage: make login EMAIL=<email>"
	@exit 1
endif
	@URL="$$(docker compose --profile cli exec -T cli go run apps/api generate-login-link $(EMAIL) | tail -n1)"; \
	if [[ "$$OSTYPE" == "linux-gnu"* ]]; then \
		xdg-open "$$URL"; \
	elif [[ "$$OSTYPE" == "darwin"* ]]; then \
		open "$$URL"; \
	else \
		echo "Unsupported OS. Please open the URL manually: $$URL"; \
	fi