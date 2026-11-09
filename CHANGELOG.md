# Changelog

## [1.1.0]

### Added

- Added weekly Dependabot updates for Go modules and GitHub Actions, plus CODEOWNERS.
- Added repository guidance and OpenSpec baseline specifications.

### Changed

- Replaced Travis CI with GitHub Actions for Go tests, builds, and Coveralls reporting on pushes to `master`.
- Updated the Go version to 1.23 and `go-runewidth` to 0.0.31; refreshed README badges and `master` references.
- Refactored tests to use `testify` against the public package API, converted example benchmark data to typed records, and replaced regular-expression ANSI CSI stripping with a byte-wise scanner.

## [1.0.0]

### Added

- Terminal tables with headers, body rows, footers, seven built-in styles, custom styles, cell alignment, multiline content, and column spanning.
- Unicode-aware display-width calculation and ANSI escape-sequence handling.
- Examples, unit tests, and benchmarks.
- Go module metadata and automated test and coverage reporting with Travis CI and Coveralls.

### Fixed

- Corrected cell sizing, padding, dividers, and alignment for spanned and Unicode text.
