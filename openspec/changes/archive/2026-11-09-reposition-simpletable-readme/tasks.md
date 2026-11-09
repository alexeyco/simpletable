## 1. Add the reusable package banner

- [x] 1.1 Update `assets/banner.svg` as a wide, asymmetrical broken-grid and row motif that suggests tabular structure without depicting a rendered table. Keep the package-name wordmark, package-specific caption, and accent replaceable; set the simpletable caption to `Simple tables in terminal with Go` and use blue for simpletable. Map to scenario `Reuse the banner for a Go package`.

## 2. Restructure the README and refresh comparisons

- [x] 2.1 Update `README.md` to start with the heading `# Simple tables in terminal with Go`, then the banner and badges, followed by a concise description, a feature matrix for simpletable, termtables, uitable, tablewriter, and go-pretty, then a short, copy-pasteable Go example using the exported API. Use `✅ documented`, `⚠️ partial`, and `❌ not found in sources checked` with a legend explaining that a cross describes the checked sources, not proven feature absence. Link the relevant public sources checked for comparison claims and date the comparison. Correct API examples to use `Cell.Text`, report seven built-in styles, and replace the full license text with a license link. Preserve the existing ASCII table output and make the Go example produce that exact output. Map to scenarios `Read the README from the top`, `Check API examples and style count`, `Review comparison claims`, `Extend the comparison`, and `Match the example output`.
- [x] 2.2 Run `go test -bench=. -benchmem` from `_example/03-benchmarks-with-others/bench` on the user's MacBook. Replace historical benchmark values in `README.md` with the observed results under a collapsed `<details>` section labeled `(MacBook)`; keep the existing comparator set. Update `_example/03-benchmarks-with-others/README.md` to retain the benchmark invocation instructions, remove the historical result block, and link to the current results in `README.md`. Map to scenarios `Expand benchmark results` and `Read the README without expanding benchmarks`.

## 3. Schedule comparison refresh reminders

- [x] 3.1 Add `.github/workflows/reminder.yml` with a monthly schedule and issue-write permission. Put the logic to create one comparison-refresh issue only when no matching issue is open in `scripts/reminder.js`, and have the workflow invoke it. Map to scenarios `Create the monthly reminder` and `Avoid duplicate open reminders`.
