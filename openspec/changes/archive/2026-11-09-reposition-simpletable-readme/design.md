# Design

## Approach

Create a standalone wide SVG banner in `assets/banner.svg`. Use an asymmetrical broken-grid and row motif to suggest tabular structure without depicting a rendered table. Keep the package-name wordmark, package-specific caption, and accent replaceable; apply the accent to selected strokes or cells. The simpletable version uses blue and displays `Simple tables in terminal with Go`. The README starts with the heading `# Simple tables in terminal with Go`, then references the banner before its badges, followed by the package summary, an evidence-linked feature matrix, and a compact example. Put detailed benchmark output inside an HTML `<details>` element so it remains available without interrupting the primary reading path.

Keep the feature comparison as a Markdown table in `README.md`. Use a fixed legend with `✅ documented`, `⚠️ partial`, and `❌ not found in sources checked`. The cross label describes the public sources checked and does not assert that a feature is absent. Link each library and comparison claim to the relevant public sources checked and include a last-checked date. Add new libraries as columns and new capabilities as rows. The initial libraries are simpletable, termtables, uitable, tablewriter, and go-pretty.

Rerun the existing benchmark suite with `go test -bench=. -benchmem` from `_example/03-benchmarks-with-others/bench` on the user's MacBook. Replace the historical output rather than relabeling the Windows measurements. Keep the benchmark harness and its comparator set unchanged; put the refreshed results in the root README's collapsible section. Keep the example README as the reproducible invocation guide and link it to the results in the root README.

Add `.github/workflows/reminder.yml` with a monthly schedule. Keep the issue lookup and creation logic in `scripts/reminder.js`, invoked by the workflow. The script checks for an already-open comparison-refresh issue and only creates one when none exists. Grant only the issue-write permission required to create the reminder. Keep the README's displayed Quick start output identical to the output produced by its code.

## Files to touch

- `assets/banner.svg` — reusable wide banner with an asymmetrical broken-grid and row motif, replaceable package-name wordmark, and accent on selected strokes or cells; blue for simpletable.
- `README.md` — opening structure, evidence-linked comparison, short example that produces its displayed ASCII table, corrected feature count, collapsible benchmark results, and license link.
- `_example/03-benchmarks-with-others/README.md` — benchmark invocation instructions and a link to the current results in the root README.
- `.github/workflows/reminder.yml` — monthly comparison-refresh issue workflow.
- `scripts/reminder.js` — comparison-refresh issue lookup and creation logic invoked by the workflow.

## Trade-offs

- A Markdown feature matrix is easy to extend and review in diffs; evidence links and a checked date make claims auditable without introducing a data-generation tool.
- The detailed benchmark output stays available but collapsed, so performance evidence does not dominate the README introduction.
- The comparison refresh is a scheduled issue rather than automated scraping; maintainers verify public evidence before changing claims.
