.PHONY: create-migration
create-migration:
	goose -dir ./migrations create ${MIGRATION_NAME} go

.PHONY: start-debug-containers
start-debug-containers:
	docker-compose -f ./deploys/debug/docker-compose.yaml up --build

.PHONY: stop-debug-containers
stop-debug-containers:
	docker-compose -f ./deploys/debug/docker-compose.yaml down

.PHONY: swagger
swagger:
	swag fmt && \
	swag init -g cmd/app/main.go --parseInternal --md docs/

.PHONY: markdown-docs
markdown-docs:
	swagger generate markdown -f docs/swagger.json --with-expand --output docs/docs.md

.PHONY: docs
docs: swagger markdown-docs