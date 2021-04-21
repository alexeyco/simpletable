all:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo

.PHONY: install
install: ## Install dependencies
	@go mod tidy \
		&& go mod vendor \
		&& go mod verify

.PHONY: lint
lint: ## Run linter
	@golangci-lint --exclude-use-default=false run ./...

runTests:
	@go test -v -coverprofile=coverage.out ./...

.PHONY: test
test: runTests ## Run tests
	@if [ -f coverage.out ]; then go tool cover -func=coverage.out && rm coverage.out; fi

.PHONY: coverage
coverage: runTests ## Show coverage
	@if [ -f coverage.out ]; then go tool cover -html=coverage.out && rm coverage.out; fi
