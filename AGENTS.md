# Agent guide

Keep this file focused on agent-specific constraints and repository gotchas. Put user-facing usage in `README.md`, contribution guidance in `CONTRIBUTING.md`, and observable product behavior in OpenSpec.

## Working in this repository

- Inspect the working tree before editing. Preserve existing user changes and keep each task's diff limited to its scope.
- Treat the Go source and tests as the authority for current behavior. Verify claims and public API details before changing documentation; do not copy potentially stale README statements as facts.
- Baseline specs in `openspec/specs/` describe behavior that exists today. Put proposed behavior in `openspec/changes/` and validate OpenSpec artifacts with `openspec validate`.
- Treat exported Go APIs and rendered table output as compatibility contracts. Make behavior changes intentional, cover them with focused tests, and update relevant documentation.
- Avoid adding dependencies when the standard library or existing packages suffice. For a new or updated dependency, check Go-version compatibility and explain the maintenance or functionality it provides.

## Go checks

- The GitHub Actions workflow runs `go test ./...` and `go build ./...`. Run the relevant checks after Go code changes; run both before reporting a code change as verified.
- The coverage workflow runs `go test -coverprofile=coverage.out ./...` and uploads its report to Coveralls on pushes to `v1`.
- For performance changes, measure a representative baseline and compare with `go test -bench . -benchmem`. Do not claim an optimization without benchmark evidence.

## Git boundaries

- Do not commit, push, merge, tag, or rewrite history unless the user explicitly requests that action.
- Do not stage, discard, or overwrite unrelated changes. Before any requested Git write, review the exact paths and preserve the rest of the working tree.
- Report checks that actually ran and their observed results; distinguish unverified behavior from verified behavior.
