include .env
MIGRATIONS_PATH=./cmd/migrate/migrations

.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS)) -source=file://$(MIGRATIONS_PATH) -database=$(DB_MIGRATIONS_ADDR) up

.PHONY: migrate-up
migrate-up:
	@migrate -source=file://$(MIGRATIONS_PATH) -database=$(DB_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -source=file://$(MIGRATIONS_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: docker-up
docker-up:
	docker compose up -d

.PHONY: docker-stop
docker-stop:
	docker compose stop

.PHONY: docker-down
docker-down:
	docker compose down
