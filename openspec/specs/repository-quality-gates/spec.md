# Repository Quality Gates Specification

## Purpose

Defines the Go test, lint, and benchmark checks run by the repository's GitHub Actions workflow and by the local `make check` aggregate.

## Requirements

### Requirement: Pull request and tag test gate

The `Quality gates` workflow SHALL run Go tests and build the root module for every pull request and version-tag push.

#### Scenario: Test and build run for supported events

**WHEN** a pull request or version tag push starts the `Quality gates` workflow
**THEN** the `Test` job runs `make test` and `make build`, which run `go test` and `go build` over the root module and the example packages in that module.

### Requirement: Lint policy

The repository SHALL configure golangci-lint with an explicit linter inventory and SHALL check `gofmt` and `goimports` formatting.

#### Scenario: Explicit linter inventory and formatters

**WHEN** the lint policy is applied to the repository
**THEN** it enables `asasalint`, `asciicheck`, `bidichk`, `copyloopvar`, `durationcheck`, `dupword`, `errcheck`, `errname`, `errorlint`, `forcetypeassert`, `gocognit`, `gocritic`, `gocyclo`, `godot`, `godox`, `gosec`, `govet`, `ineffassign`, `makezero`, `misspell`, `nilerr`, `nilnil`, `nolintlint`, `predeclared`, `revive`, `staticcheck`, `testifylint`, `testpackage`, `thelper`, `unconvert`, `unparam`, `unused`, `usetesting`, and `wastedassign`; checks `gofmt` and `goimports`; and does not use the moving `all` set.

### Requirement: Pull request and tag lint gate

The `Quality gates` workflow SHALL run a blocking `Lint` job for every pull request and version-tag push. The job SHALL install `golangci-lint` without running it directly, then execute `make lint` and `make fmt-check`, which lint the root Go module and all six Go example packages and fail when lint or formatting findings remain.

#### Scenario: Lint covers the library and examples

**WHEN** a pull request or version tag push starts the `Quality gates` workflow
**THEN** the `Lint` job analyzes the root module, the four example packages in that module, and both packages in the isolated benchmark comparison module.

#### Scenario: Lint findings block the check

**WHEN** a checked Go package has an enabled-linter or formatting finding
**THEN** the `Lint` job fails until the finding is fixed, without a temporary baseline.

### Requirement: Test assertion helpers

The repository SHALL use `testify` assertions in unit tests where they replace repeated manual failure checks, and SHALL enable `testifylint` to check assertion usage.

#### Scenario: Testify assertions are linted

**WHEN** the `Lint` job analyzes repository tests
**THEN** `testifylint` checks assertion usage and repeated output and length checks use `testify/assert` or `testify/require` helpers.

### Requirement: Print API compatibility during linting

The lint policy SHALL preserve the public signatures of `Table.Print` and `Table.Println` and SHALL exclude only `fmt.Print` and `fmt.Println` from `errcheck` for those methods.

#### Scenario: Print methods retain their signatures

**WHEN** lint findings in `Table.Print` or `Table.Println` are addressed
**THEN** both methods retain their existing no-return signatures and `errcheck` exempts only the `fmt.Print` and `fmt.Println` calls.

### Requirement: Comparison benchmark gate

`make bench-compare` SHALL run the comparison benchmarks in the `bench` package of `_example/03-benchmarks-with-others`, reporting `BenchmarkSimpletable`, `BenchmarkTermtables`, and `BenchmarkUITable` followed by `PASS`. `make bench-compare-check` SHALL fail unless that package's benchmark listing contains all three names, SHALL be a prerequisite of `make check`, and SHALL appear in `make help`.

#### Scenario: bench-compare runs the three named benchmarks

**WHEN** `make bench-compare` runs from the repository root
**THEN** its output reports results for `BenchmarkSimpletable`, `BenchmarkTermtables`, and `BenchmarkUITable` followed by `PASS`

#### Scenario: bench-compare-check passes on the current tree

**WHEN** `make bench-compare-check` runs on a tree whose `bench` package defines the three comparison benchmarks
**THEN** the listing of that package contains all three names and the target exits successfully

#### Scenario: bench-compare-check fails when a name is missing

**WHEN** the `bench` package listing lacks any of `BenchmarkSimpletable`, `BenchmarkTermtables`, or `BenchmarkUITable`
**THEN** `make bench-compare-check` exits non-zero and names the missing benchmark

#### Scenario: make check includes the lock and stays green

**WHEN** `make check` runs on the fixed tree
**THEN** it executes `bench-compare-check` alongside the existing `test`, `build`, `lint`, `fmt-check`, and `openspec-validate` prerequisites and exits successfully

#### Scenario: make help lists the new target

**WHEN** `make help` runs
**THEN** it lists `bench-compare-check` with its description in addition to every existing target

### Requirement: Comparison benchmark documentation

The `make help` description of `bench-compare`, `AGENTS.md`, and the harness README `## Run benchmarks` section SHALL point at the `bench` package. The harness README `## Installation` section SHALL describe a module-aware way to obtain the comparators pinned in that module's `go.mod` — `go mod download` from the module directory — and SHALL contain no `go get` line. The `## Run benchmarks` heading SHALL stay byte-identical so the root README's `#run-benchmarks` anchor keeps resolving.

#### Scenario: Documentation surfaces agree on the bench package

**WHEN** a reader consults `AGENTS.md`, the `make help` description of `bench-compare`, or the harness README `## Run benchmarks` section
**THEN** each points at the `bench` package of the comparison module, the section also names `make bench-compare` as the repo-root invocation, and the `## Run benchmarks` heading is unchanged so the root README's `#run-benchmarks` anchor keeps resolving

#### Scenario: Installation section describes module-aware retrieval

**WHEN** a reader obtains the comparators by following the harness README `## Installation` section
**THEN** the section states that the comparators are pinned in the module's `go.mod` and fetched automatically, names `go mod download` run from the module directory as the explicit retrieval step, and contains no `go get` line
