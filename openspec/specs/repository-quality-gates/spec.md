# Repository Quality Gates Specification

## Purpose

Defines the Go test and lint checks run by the repository's GitHub Actions workflow.

## Requirements

### Requirement: Pull request and tag test gate

The `Quality gates` workflow SHALL run Go tests and build the root module for every pull request and version-tag push.

#### Scenario: Test and build run for supported events

**WHEN** a pull request or version tag push starts the `Quality gates` workflow
**THEN** the `Test` job runs `go test ./...` and `go build ./...`.

### Requirement: Lint policy

The repository SHALL configure golangci-lint with an explicit linter inventory and SHALL check `gofmt` and `goimports` formatting.

#### Scenario: Explicit linter inventory and formatters

**WHEN** the lint policy is applied to the repository
**THEN** it enables `asasalint`, `asciicheck`, `bidichk`, `copyloopvar`, `durationcheck`, `dupword`, `errcheck`, `errname`, `errorlint`, `forcetypeassert`, `gocognit`, `gocritic`, `gocyclo`, `godot`, `godox`, `gosec`, `govet`, `ineffassign`, `makezero`, `misspell`, `nilerr`, `nilnil`, `nolintlint`, `predeclared`, `revive`, `staticcheck`, `testifylint`, `testpackage`, `thelper`, `unconvert`, `unparam`, `unused`, `usetesting`, and `wastedassign`; checks `gofmt` and `goimports`; and does not use the moving `all` set.

### Requirement: Pull request and tag lint gate

The `Quality gates` workflow SHALL run a blocking `Lint` job for every pull request and version-tag push. The job SHALL lint the root Go module and all six Go example packages and fail when lint or formatting findings remain.

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
