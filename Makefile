current_dir := $(shell pwd)

.PHONY: gen-docs
gen-docs:
	docker run --rm -v "$(current_dir):/code" \
		ghcr.io/swaggo/swag:v1.16.4 init \
		-g ./api/main.go \
		-d cmd,internal \
		-o internal/docs
	docker run --rm -v "$(current_dir):/code" \
		ghcr.io/swaggo/swag:v1.16.4 fmt \
		-d internal/docs
	docker run --rm -v "$(current_dir):/code" \
		ghcr.io/swaggo/swag:v1.16.4 fmt \
		-d .

.PHONY: reset-all
reset-all:
	docker compose down -v
	docker compose up -d

.PHONY: regenerate-db
regenerate-db:
	docker compose stop db
	docker compose rm -f -v db
	docker compose up -d db

.PHONY: seed
seed: regenerate-db
	docker compose exec -T db bash -c "until pg_isready -U $${POSTGRES_USER:-postgres} -d $${POSTGRES_DB:-paw-me-back}; do sleep 1; done"
	@go run cmd/seed/main.go