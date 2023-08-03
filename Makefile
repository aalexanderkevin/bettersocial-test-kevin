dep:
	@echo ">> Downloading Dependencies"
	@go mod download

generate-rest-api-docs: dep
	@echo ">> Generating REST API docs with Swagger"
	@swag init -dir ./controller -g routes.go -o ./docs/api/rest/swag --parseDependency --parseInternal --parseDepth 3

run-server: dep
	env $$(cat .env | xargs) go run bettersocial/cmd server

test-all: test-unit test-integration-with-infra

test-unit: dep
	@echo ">> Running Unit Test"
	@env $$(cat .env.testing | xargs) go test -tags=unit -failfast -cover -covermode=atomic ./...

test-integration: dep
	@echo ">> Running Integration Test"
	@env $$(cat .env.testing | xargs) env POSTGRES_MIGRATION_PATH=$$(pwd)/database/migrations go test -tags=integration -failfast -cover -covermode=atomic ./...

test-integration-with-infra: test-infra-up test-integration test-infra-down

test-infra-up:
	$(MAKE) test-infra-down
	@echo ">> Starting Test DB"
	docker run -d --rm --name test-postgres -p 5431:5432 --env-file .env.testing postgres:12

test-infra-down:
	@echo ">> Shutting Down Test DB"
	@-docker kill test-postgres

migrate:
	eval $$(egrep -v '^#' .env | xargs -0) go run bettersocial/cmd migrate