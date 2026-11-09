# Design

## Approach

Use a root `.golangci.yml` in golangci-lint v2 format and pin the matching golangci-lint v2 release in the `Quality gates` workflow. Enable the 34 confirmed linters explicitly, including `testifylint`, configure `gofmt` and `goimports`, and keep `errcheck` exclusions limited to `fmt.Print` and `fmt.Println`. Do not enable `gochecknoglobals`; retain exported style variables and typed package-level benchmark fixtures. Run `govet` through golangci-lint; do not create a duplicate `go vet` job.

The `Quality gates` workflow runs on pull requests, version tags, and weekly on Monday. It keeps test/build jobs off scheduled runs, and runs `Lint` only on pull requests. It runs `Vulnerability` on pull requests and weekly. Lint checks the root module, the four example packages that remain in that module, and both packages in the separately isolated benchmark comparison example module. Run `golangci-lint fmt --diff` and the configured lint checks from both module roots; from `_example/03-benchmarks-with-others`, pass the root configuration with `--config ../../.golangci.yml`. Run pinned `actionlint` on workflow files. Each example path under the root module is included explicitly because wildcard patterns do not include `_example`: `./_example/01-styles-demo`, `./_example/02-ugly-span`, `./_example/04-multiline`, and `./_example/05-colorize`.

The `Vulnerability` job installs govulncheck v1.1.4, which is compatible with Go 1.23, and uses text output so reachable vulnerability findings produce a failing check. It scans the root module, the four example packages that remain in that module, and both packages in the isolated benchmark comparison example module. The `Quality gates` workflow uses read-only contents permission. The `Coveralls` workflow grants the status-write permission required for coverage publication.

Use the exact `Lint` and `Vulnerability` job names as required checks in the `master` branch protection settings. The checks remain distinct so the merge gate reports lint and vulnerability results separately.

## Code changes

Keep the exported style variables in `style.go` for compatibility. Keep benchmark fixtures as typed package-level variables so renderer calls reuse the same rows without rebuilding a slice on every call. Move tests in `benchmarks_test.go`, `cell_test.go`, `column_test.go`, `content_test.go`, `row_test.go`, and `table_test.go` to package `simpletable_test`; express coverage through exported behavior while preserving the existing output and API checks. Move `_example/03-benchmarks-with-others/bench/simpletable_test.go`, `_example/03-benchmarks-with-others/bench/termtables_test.go`, and `_example/03-benchmarks-with-others/bench/uitable_test.go` to package `bench_test`. Keep the `Table.Print` and `Table.Println` declarations in `table.go` unchanged.

Add a module at `_example/03-benchmarks-with-others/go.mod` and its generated `_example/03-benchmarks-with-others/go.sum`. This module owns the example-only `termtables` and `uitable` dependencies, includes the example `main` package and comparator package, and uses the root library through a local module replacement. This keeps those comparator dependencies out of the library's `go.mod` and `go.sum` while preserving a buildable example.

Apply all reported findings in the linted files: `cell.go`, `cell_test.go`, `column.go`, `column_test.go`, `content.go`, `content_test.go`, `contents.go`, `doc.go`, `row.go`, `row_test.go`, `style.go`, `table.go`, `table_test.go`, `benchmarks_test.go`, `_example/01-styles-demo/main.go`, `_example/02-ugly-span/main.go`, `_example/03-benchmarks-with-others/main.go`, `_example/03-benchmarks-with-others/bench/data.go`, `_example/03-benchmarks-with-others/bench/simpletable.go`, `_example/03-benchmarks-with-others/bench/simpletable_test.go`, `_example/03-benchmarks-with-others/bench/termtables.go`, `_example/03-benchmarks-with-others/bench/termtables_test.go`, `_example/03-benchmarks-with-others/bench/uitable.go`, `_example/03-benchmarks-with-others/bench/uitable_test.go`, `_example/04-multiline/main.go`, and `_example/05-colorize/main.go`.

## Files and settings

- `.golangci.yml` — explicit linter inventory, formatter configuration, and narrow `errcheck` exclusions.
- `.github/workflows/quality-gates.yml` — test/build jobs, PR `Lint` job, and PR/weekly `Vulnerability` job; pins golangci-lint v2, actionlint, and govulncheck v1.1.4.
- `.github/workflows/coveralls.yml` — coverage upload on pushes to `master`.
- `style.go`, `benchmarks_test.go`, and the Go source/test files listed above — lint remediations while preserving public output and API contracts.
- `_example/03-benchmarks-with-others/go.mod` and `_example/03-benchmarks-with-others/go.sum` — isolated benchmark comparison example module and dependency checksums.
- `go.mod` and `go.sum` — `testify` test assertions and their checksums.
- GitHub branch protection settings for `master` — required checks `Lint` and `Vulnerability`.

## Trade-offs

- Pinning govulncheck v1.1.4 keeps the scanner compatible with Go 1.23, while using an older scanner implementation than current releases. Its advisory data remains queried from the live Go vulnerability database.
- Explicit linter names make upgrades predictable and avoid new linters entering the gate implicitly.
- The explicit lint inventory omits `gochecknoglobals`, preserving exported style compatibility values and reusable typed benchmark fixtures without local suppressions.
- The benchmark comparison example module adds a module boundary so its dependencies do not become library dependencies.
