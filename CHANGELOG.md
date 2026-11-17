# Changelog

## 1.2.1

### Added

- Added frozen-output golden tests (`testdata/*.golden`) that pin `Table.String()` byte for byte across all seven built-in styles and multiline, ANSI-colored, CJK-wide, and column-spanned content, plus repeated-render and concurrent-render checks of the recycled cell content.

### Changed

- Reduced per-render allocations in `Table.String()` by about 50 percent, measured on the root `BenchmarkStyle*` suite, by building rendered lines into reused buffers, replacing per-line `fmt.Sprintf` and `strings.Repeat` calls, pre-sizing slices, and pooling internal cell content objects; rendered output and the exported API are unchanged.
- Refreshed the README benchmark figures from the comparison harness after the allocation reductions.

## 1.2.0

### Added

- Added a `Makefile` defining all quality-gate commands (`make help` lists targets; `make check` runs tests, build, lint, format checks, and OpenSpec validation) and pinned agent skill management (`make skills`, `make skills-update`).
- Added project-level agent skills pinned in `skills-lock.json`: Go code style, testing, performance, and error-handling guidelines, plus a two-axis code-review workflow.
- Added a monthly GitHub Actions workflow that opens an issue reminding maintainers to review and refresh the README comparison.

### Changed

- Replaced inline commands in the `Quality gates` and `Coveralls` workflows with calls to the shared Makefile targets, making the `Makefile` the single source for gate commands; the `Test` job now also builds and tests the example packages.
- Repositioned the README around a reusable banner, a concise package summary, a feature comparison with an evidence legend, and a short quick-start example; corrected examples and feature counts to match the package API.

### Fixed

- Fixed broken row alignment when a spanning cell's content widened shared columns after another spanning cell had already taken its width from them; column growth and spanning-cell width syncing now run as separate passes, so rendered geometry no longer depends on the order of spanning cells (#23).

## 1.1.0

### Added

- Added weekly Dependabot updates for Go modules and GitHub Actions, plus CODEOWNERS.
- Added repository guidance and OpenSpec baseline specifications.

### Changed

- Replaced Travis CI with GitHub Actions for Go tests, builds, and Coveralls reporting on pushes to `master`.
- Updated the Go version to 1.23 and `go-runewidth` to 0.0.31; refreshed README badges and `master` references.
- Refactored tests to use `testify` against the public package API, converted example benchmark data to typed records, and replaced regular-expression ANSI CSI stripping with a byte-wise scanner.

## 1.0.0

### Added

- Terminal tables with headers, body rows, footers, seven built-in styles, custom styles, cell alignment, multiline content, and column spanning.
- Unicode-aware display-width calculation and ANSI escape-sequence handling.
- Examples, unit tests, and benchmarks.
- Go module metadata and automated test and coverage reporting with Travis CI and Coveralls.

### Fixed

- Corrected cell sizing, padding, dividers, and alignment for spanned and Unicode text.
