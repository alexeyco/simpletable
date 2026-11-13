.DEFAULT_GOAL := help

# The Go toolchain skips directories beginning with "_", so ./... never
# covers the example packages; they must be listed explicitly.
EXAMPLES := ./_example/01-styles-demo ./_example/02-ugly-span ./_example/04-multiline ./_example/05-colorize
BENCH_DIR := _example/03-benchmarks-with-others

## Help

.PHONY: help
help: ## List available targets
	@awk 'BEGIN {FS = ":.*## "} /^[A-Za-z0-9_-]+:.*## / {printf "  %-18s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

## Go checks

.PHONY: test
test: ## Run tests for the root module and the examples
	go test ./... $(EXAMPLES)

.PHONY: build
build: ## Build the root module and the examples
	go build ./... $(EXAMPLES)

.PHONY: lint
lint: ## Run golangci-lint over every module, as CI does
	golangci-lint run ./... $(EXAMPLES)
	cd $(BENCH_DIR) && golangci-lint run --config=../../.golangci.yml ./...

.PHONY: fmt
fmt: ## Format all modules with golangci-lint
	golangci-lint fmt --config .golangci.yml ./... $(EXAMPLES)
	cd $(BENCH_DIR) && golangci-lint fmt --config ../../.golangci.yml

.PHONY: fmt-check
fmt-check: ## Fail if any module is not formatted, as CI does
	golangci-lint fmt --config .golangci.yml --diff ./... $(EXAMPLES)
	cd $(BENCH_DIR) && golangci-lint fmt --config ../../.golangci.yml --diff

.PHONY: cover
cover: ## Produce coverage.out, the same file the Coveralls action uploads
	go test -coverprofile=coverage.out ./...

## Benchmarks

.PHONY: bench
bench: ## Run root-module benchmarks with allocation counts
	go test -run '^$$' -bench . -benchmem .

.PHONY: bench-compare
bench-compare: ## Run third-party comparison benchmarks (module 03)
	cd $(BENCH_DIR) && go test -bench . -benchmem

## OpenSpec

.PHONY: openspec-validate
openspec-validate: ## Validate all OpenSpec changes and specs
	openspec validate --all --strict --no-interactive

## Aggregates

.PHONY: check
check: test build lint fmt-check openspec-validate ## Run every local gate: test, build, lint, format, OpenSpec

## Skills

.PHONY: skills
skills: ## Restore pinned agent skills from skills-lock.json into .agents/skills
	# "experimental_install" is the upstream CLI name for restoring exactly
	# what skills-lock.json pins; it is marked experimental by the skills CLI.
	npx -y skills experimental_install

.PHONY: skills-update
skills-update: ## Refresh skills from upstream and update skills-lock.json
	npx -y skills update -p -y
