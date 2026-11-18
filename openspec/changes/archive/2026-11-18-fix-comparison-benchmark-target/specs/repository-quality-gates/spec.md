## ADDED Requirements

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
