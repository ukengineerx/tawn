.PHONY: dev-up dev-down run-api test check web-dev mobile-start

COMPOSE = docker compose -f deployments/local/docker-compose.yml

dev-up:
	$(COMPOSE) up -d --wait

dev-down:
	$(COMPOSE) down

run-api:
	cd backend && go run ./cmd/api

test:
	cd backend && go test ./...

check: test
	cd backend && go vet ./...
	cd web && npm run type-check && npm run lint
	cd mobile/app && npm run typecheck

web-dev:
	cd web && npm run dev

mobile-start:
	cd mobile/app && npm start
