# Add Go quality gates

## Why now

The repository already tests and builds Go code on pull requests, and Dependabot checks Go modules and GitHub Actions weekly. It has no blocking lint check or vulnerability scan, so those checks cannot catch code-quality findings or newly published advisories before or after merge.

## What changes

Add a pinned golangci-lint v2 configuration with an explicit, broad linter inventory, formatting checks, and pinned `actionlint`. Resolve the findings instead of keeping a baseline. Add a `Lint` job to the `Quality gates` workflow for pull requests; cover the root module and all six Go example packages. Add `testify` assertions to shorten repeated manual checks and enable `testifylint`. Preserve the `Table.Print` and `Table.Println` signatures while narrowly excluding their `fmt.Print` and `fmt.Println` calls from `errcheck`.

Add a `Vulnerability` job to the same `Quality gates` workflow using govulncheck v1.1.4 on pull requests and weekly. Fail only when govulncheck finds a vulnerability reachable from code. Require both checks in the `master` branch protection settings. Keep the existing Dependabot schedule and Go test/build jobs; do not add a separate `go vet ./...` job because `govet` is part of the lint set.

## Impact on capabilities

Add the `repository-quality-gates` capability for linting, vulnerability scanning, and merge requirements. Existing table rendering, layout, and style capabilities do not change.

DECISION: Preserve exported built-in `Style*` variables and typed package-level benchmark fixtures. Omit `gochecknoglobals` from the explicit lint inventory so renderers can reuse fixture data without allocating a fresh slice on every call.

ASSUMPTION: Put the entire benchmark comparison example, including its `main` package and comparator package, in a separate Go module so its two example-only dependencies do not enter the library module and the example remains buildable.

ASSUMPTION: Convert tests to the external `simpletable_test` package for testpackage while preserving observable behavior coverage.

REJECTED-RISK: Use a current govulncheck requiring Go 1.25/1.26 — user chose a pinned scanner compatible with the repository's Go 1.23 declaration.

REJECTED-RISK: Block on any vulnerable module version, including unreachable code — user chose to block only on govulncheck vulnerabilities reachable from application code.

REJECTED-RISK: Run a standalone `go vet ./...` step — user chose `govet` inside golangci-lint as sufficient.

REJECTED-RISK: Allow existing lint findings through a temporary baseline — user chose blocking lint checks and fixing findings.

REJECTED-RISK: Enable a moving `all` linter set — user chose a broad, pinned explicit list for reproducible CI.
