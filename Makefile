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


.PHONY: regenerate-db
regenerate-db:
	docker compose down -v
	docker compose up -d

.PHONY: seed
seed: regenerate-db
	@go run cmd/seed/main.go
