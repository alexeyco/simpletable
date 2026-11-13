# Agent guide

Keep this file focused on agent-specific constraints and repository gotchas. Put user-facing usage in `README.md`, contribution guidance in `CONTRIBUTING.md`, and observable product behavior in OpenSpec.

## Working in this repository

- Inspect the working tree before editing. Preserve existing user changes and keep each task's diff limited to its scope.
- Treat the Go source and tests as the authority for current behavior. Verify claims and public API details before changing documentation; do not copy potentially stale README statements as facts.
- Baseline specs in `openspec/specs/` describe behavior that exists today. Put proposed behavior in `openspec/changes/` and validate OpenSpec artifacts with `openspec validate`.
- Treat exported Go APIs and rendered table output as compatibility contracts. Make behavior changes intentional, cover them with focused tests, and update relevant documentation.
- Avoid adding dependencies when the standard library or existing packages suffice. For a new or updated dependency, check Go-version compatibility and explain the maintenance or functionality it provides.

## Go checks

- Run checks through the `Makefile`; the CI workflows call these same targets, so the Makefile is the single source for gate commands. Use `make help` to list targets. Run `make check` before reporting a code change as verified; run the narrowest relevant target (`make test`, `make lint`, `make fmt-check`, `make openspec-validate`) while iterating.
- The coverage workflow runs `make cover` (`go test -coverprofile=coverage.out ./...`) and uploads its report to Coveralls on pushes to `master`.
- For performance changes, measure a representative baseline with `make bench` and compare afterwards. Do not claim an optimization without benchmark evidence. `make bench-compare` runs the third-party comparison benchmarks in `_example/03-benchmarks-with-others`.

## Agent skills

- Project skills are pinned in `skills-lock.json` and installed into `.agents/skills/`, which is gitignored. Run `make skills` to restore them after cloning; run `make skills-update` to refresh them from upstream.

## Git boundaries

- Do not commit, push, merge, tag, or rewrite history unless the user explicitly requests that action.
- Do not stage, discard, or overwrite unrelated changes. Before any requested Git write, review the exact paths and preserve the rest of the working tree.
- Report checks that actually ran and their observed results; distinguish unverified behavior from verified behavior.
